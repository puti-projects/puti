package page

import (
	"fmt"

	Response "github.com/puti-projects/puti/internal/backend/handler"
	"github.com/puti-projects/puti/internal/common/model"
	"github.com/puti-projects/puti/internal/common/utils"
	"github.com/puti-projects/puti/internal/pkg/errno"
	"github.com/puti-projects/puti/internal/pkg/token"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// CreateRequest struct of page create params
type CreateRequest struct {
	Status        string `json:"status"`
	Title         string `json:"title"`
	Content       string `json:"content"`
	ContentHTML   string `json:"content_html"`
	Description   string `json:"description"`
	CommentStatus uint64 `json:"comment_status"`
	CoverPicture  string `json:"cover_picture"`
	PostedTime    string `json:"posted_time"`
	Slug          string `json:"slug"`
	PageTemplate  string `json:"page_template"`
	ParentID      uint64 `json:"parent_id"`
}

// CreateResponse return the new page id and url
type CreateResponse struct {
	ID   uint64 `json:"id"`
	GUID string `json:"guid"`
}

// Create add new page
func Create(c *gin.Context) {
	log.Info("Page create function called.", lager.Data{"X-request-Id": utils.GetReqID(c)})

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
		Response.SendResponse(c, errno.ErrPageCreateFailed, nil)
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
	page := model.PostModel{
		UserID:          userID,
		PostType:        "page",
		Title:           r.Title,
		ContentMarkdown: r.Content,
		ContentHTML:     r.ContentHTML,
		Slug:            r.Slug,
		ParentID:        r.ParentID,
		Status:          r.Status,
		CommentStatus:   r.CommentStatus,
		IfTop:           0,
		GUID:            fmt.Sprintf("/%s", r.Slug),
		CoverPicture:    r.CoverPicture,
		CommentCount:    0,
		ViewCount:       0,
		PostDate:        utils.StringToTime("2006-01-02 15:04:05", r.PostedTime),
	}
	if err := tx.Create(&page).Error; err != nil {
		return rsp, err
	}

	// set metadata description
	metaDescription := model.PostMetaModel{
		PostID:    page.ID,
		MetaKey:   "description",
		MetaValue: r.Description,
	}
	if err := tx.Create(&metaDescription).Error; err != nil {
		tx.Rollback()
		return rsp, err
	}

	// set metadata page_template
	metaPageTemplate := model.PostMetaModel{
		PostID:    page.ID,
		MetaKey:   "page_template",
		MetaValue: r.PageTemplate,
	}
	if err := tx.Create(&metaPageTemplate).Error; err != nil {
		tx.Rollback()
		return rsp, err
	}

	rsp.ID = page.ID
	rsp.GUID = page.GUID

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

	if isExist := model.PageCheckSlugExist(0, r.Slug); isExist == true {
		return errno.New(errno.ErrSlugExist, nil)
	}

	return nil
}
