package article

import (
	"fmt"
	"strconv"
	"strings"

	Response "puti/handler"
	"puti/model"
	"puti/pkg/errno"
	"puti/pkg/token"
	"puti/service"
	"puti/utils"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// CreateRequest struct of article create params
type CreateRequest struct {
	Status        string   `json:"status"`
	Title         string   `json:"title"`
	Content       string   `json:"content"`
	ContentHTML   string   `json:"content_html"`
	Description   string   `json:"description"`
	CommentStatus uint64   `json:"comment_status"`
	CoverPicture  string   `json:"cover_picture"`
	PostedTime    string   `json:"posted_time"`
	IfTop         uint64   `json:"if_top"`
	Category      []uint64 `json:"category"`
	Tag           []uint64 `json:"tag"`
}

// CreateResponse return the new article id and url
type CreateResponse struct {
	ID   uint64 `json:"id"`
	GUID string `json:"guid"`
}

// Create create a new aricle(published or draft)
func Create(c *gin.Context) {
	log.Info("Article create function called.", lager.Data{"X-request-Id": utils.GetReqID(c)})

	// get token and parse
	t := c.Query("token")
	userContext, err := token.ParseToken(t)

	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		Response.SendResponse(c, errno.ErrBind, nil)
		return
	}

	// check params
	if err := r.checkParam(); err != nil {
		Response.SendResponse(c, err, nil)
		return
	}

	// add article data
	rsp, err := handleCreate(&r, userContext.ID)
	if err != nil {
		Response.SendResponse(c, errno.ErrArticleCreateFailed, nil)
		return
	}

	Response.SendResponse(c, nil, rsp)
}

func handleCreate(r *CreateRequest, userID uint64) (rsp *CreateResponse, err error) {
	rsp = new(CreateResponse)

	tx := model.DB.Local.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// main data
	article := model.PostModel{
		UserID:          userID,
		PostType:        "article",
		Title:           r.Title,
		ContentMarkdown: r.Content,
		ContetnHTML:     r.ContentHTML,
		ParentID:        0,
		Status:          r.Status,
		CommentStatus:   r.CommentStatus,
		IfTop:           r.IfTop,
		CoverPicture:    r.CoverPicture,
		CommentCount:    0,
		ViewCount:       0,
		PostDate:        utils.StringToTime("2006-01-02 15:04:05", r.PostedTime),
	}
	if err := tx.Create(&article).Error; err != nil {
		return rsp, err
	}

	// set GUID
	article.GUID = fmt.Sprintf("/article/%s.html", strconv.FormatUint(article.ID, 10))
	if err := tx.Model(&model.PostModel{}).Save(article).Error; err != nil {
		tx.Rollback()
		return rsp, err
	}

	// set metadata
	articleMeta := model.PostMetaModel{
		PostID:    article.ID,
		MetaKey:   "description",
		MetaValue: r.Description,
	}
	if err := tx.Create(&articleMeta).Error; err != nil {
		tx.Rollback()
		return rsp, err
	}

	// set category and tag
	valueStrings := make([]string, 0, len(r.Category)+len(r.Tag))
	valueArgs := make([]interface{}, 0, (len(r.Category) + len(r.Tag)*3))
	for _, category := range r.Category {
		termTaxonomy, _ := model.GetTermTaxonomy(category, "category")
		valueStrings = append(valueStrings, "(?, ?, ?)")
		valueArgs = append(valueArgs, article.ID)      // object_id
		valueArgs = append(valueArgs, termTaxonomy.ID) // term_taxonomy_id
		valueArgs = append(valueArgs, 0)               // term_order
	}
	for _, tag := range r.Tag {
		termTaxonomy, _ := model.GetTermTaxonomy(tag, "tag")
		valueStrings = append(valueStrings, "(?, ?, ?)")
		valueArgs = append(valueArgs, article.ID)      // object_id
		valueArgs = append(valueArgs, termTaxonomy.ID) // term_taxonomy_id
		valueArgs = append(valueArgs, 0)               // term_order
	}
	tb := &model.TermRelationshipsModel{}
	stmt := fmt.Sprintf("INSERT INTO %s (object_id, term_taxonomy_id, term_order) VALUES %s", tb.TableName(), strings.Join(valueStrings, ","))
	if err := tx.Exec(stmt, valueArgs...).Error; err != nil {
		tx.Rollback()
		return rsp, err
	}

	// update taxonomy count
	insertTaxonomy := append(r.Category, r.Tag...)
	if err := service.UpdateTaxonomyCountByArticleChange(tx, insertTaxonomy, 1); err != nil {
		tx.Rollback()
		return nil, err
	}

	rsp.ID = article.ID
	rsp.GUID = article.GUID

	// commit
	return rsp, tx.Commit().Error
}

func (r *CreateRequest) checkParam() error {
	if r.Title == "" {
		return errno.New(errno.ErrValidation, nil).Add("Title can not be empty.")
	}

	if r.Content == "" {
		return errno.New(errno.ErrValidation, nil).Add("Content can not be empty.")
	}

	if r.PostedTime == "" {
		return errno.New(errno.ErrValidation, nil).Add("PostedTime can not be empty.")
	}

	if r.Status == "" {
		return errno.New(errno.ErrValidation, nil).Add("Status can not be empty.")
	}

	if r.Status != "publish" && r.Status != "draft" {
		return errno.New(errno.ErrValidation, nil).Add("Status is incorrect.")
	}

	return nil
}
