package article

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	Response "github.com/puti-projects/puti/internal/backend/handler"
	"github.com/puti-projects/puti/internal/backend/service"
	"github.com/puti-projects/puti/internal/model"
	"github.com/puti-projects/puti/internal/pkg/db"
	"github.com/puti-projects/puti/internal/pkg/errno"
	"github.com/puti-projects/puti/internal/pkg/token"
	"github.com/puti-projects/puti/internal/utils"

	"github.com/gin-gonic/gin"
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
	Subject       []uint64 `json:"subject"`
}

// CreateResponse return the new article id and url
type CreateResponse struct {
	ID   uint64 `json:"id"`
	GUID string `json:"guid"`
}

// Create create a new article (published or draft)
func Create(c *gin.Context) {
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
		fmt.Println(err)
		Response.SendResponse(c, errno.ErrArticleCreateFailed, nil)
		return
	}

	Response.SendResponse(c, nil, rsp)
}

func handleCreate(r *CreateRequest, userID uint64) (rsp *CreateResponse, err error) {
	rsp = new(CreateResponse)

	tx := db.DBEngine.Begin()
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
		ContentHTML:     r.ContentHTML,
		ParentID:        0,
		Status:          r.Status,
		CommentStatus:   r.CommentStatus,
		IfTop:           r.IfTop,
		CoverPicture:    r.CoverPicture,
		CommentCount:    0,
		ViewCount:       0,
	}
	if r.PostedTime == "" && r.Status == model.PostStatusPublish {
		article.PostDate = &sql.NullTime{Time: time.Now(), Valid: true}
	} else {
		article.PostDate = utils.StringToNullTime("2006-01-02 15:04:05", r.PostedTime)
	}
	if err := tx.Create(&article).Error; err != nil {
		return rsp, err
	}

	// set GUID
	article.GUID = fmt.Sprintf("/article/%s.html", strconv.FormatUint(uint64(article.ID), 10))
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
	valueArgs := make([]interface{}, 0, len(r.Category)+len(r.Tag)*3)
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
		return rsp, err
	}

	// if upload subject id
	if subjectLen := len(r.Subject); subjectLen != 0 {
		subjectValueStrings := make([]string, 0, subjectLen)
		subjectValueArgs := make([]interface{}, 0, subjectLen*3)
		for _, subject := range r.Subject {
			if subject != 0 {
				subjectValueStrings = append(subjectValueStrings, "(?, ?, ?)")
				subjectValueArgs = append(subjectValueArgs, article.ID)
				subjectValueArgs = append(subjectValueArgs, subject)
				subjectValueArgs = append(subjectValueArgs, 0)
			}
		}
		sr := &model.SubjectRelationshipsModel{}
		sqls := fmt.Sprintf("INSERT INTO %s (object_id, subject_id, order_num) VALUES %s", sr.TableName(), strings.Join(subjectValueStrings, ","))
		if err := tx.Exec(sqls, subjectValueArgs...).Error; err != nil {
			tx.Rollback()
			return rsp, err
		}

		// update subject
		if err := service.UpdateSubjectInfoByArticleChange(tx, r.Subject, 1, true); err != nil {
			tx.Rollback()
			return rsp, err
		}
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

	if r.Status == "" {
		return errno.New(errno.ErrValidation, nil).Add("Status can not be empty.")
	}

	if r.Status != "publish" && r.Status != "draft" {
		return errno.New(errno.ErrValidation, nil).Add("Status is incorrect.")
	}

	return nil
}
