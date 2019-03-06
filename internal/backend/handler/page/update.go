package page

import (
	"strconv"

	Response "github.com/puti-projects/puti/internal/backend/handler"
	"github.com/puti-projects/puti/internal/backend/service"
	"github.com/puti-projects/puti/internal/common/model"
	"github.com/puti-projects/puti/internal/common/utils"
	"github.com/puti-projects/puti/internal/pkg/errno"
	"github.com/puti-projects/puti/internal/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// UpdateRequest struct of page update params
type UpdateRequest struct {
	ID            uint64 `json:"id"`
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

// Update update page info
// Delete and restore info are also in this function and it depends on the 'status'
func Update(c *gin.Context) {
	logger.Info("Page update function called.", zap.String("X-request-Id", utils.GetReqID(c)))

	// Get page id
	ID, _ := strconv.Atoi(c.Param("id"))

	var r UpdateRequest
	if err := c.ShouldBind(&r); err != nil {
		Response.SendResponse(c, errno.ErrBind, nil)
		return
	}

	pageID := uint64(ID)

	// check params
	if err := r.checkParam(pageID); err != nil {
		Response.SendResponse(c, err, nil)
		return
	}

	if r.Status == "deleted" {
		if err := service.TrashPost(pageID); err != nil {
			Response.SendResponse(c, errno.ErrDatabase, nil)
			return
		}
	} else if r.Status == "restore" {
		if err := service.RestorePost(pageID); err != nil {
			Response.SendResponse(c, errno.ErrDatabase, nil)
			return
		}
	} else {
		page := &model.PostModel{
			Model:           model.Model{ID: r.ID},
			Title:           r.Title,
			ContentMarkdown: r.Content,
			ContentHTML:     r.ContentHTML,
			Status:          r.Status,
			CommentStatus:   r.CommentStatus,
			CoverPicture:    r.CoverPicture,
			PostDate:        utils.StringToNullTime("2006-01-02 15:04:05", r.PostedTime),
			Slug:            r.Slug,
			ParentID:        r.ParentID,
		}

		// Update changed fields.
		if err := service.UpdatePage(page, r.Description, r.PageTemplate); err != nil {
			Response.SendResponse(c, errno.ErrDatabase, nil)
			return
		}
	}

	Response.SendResponse(c, nil, nil)
	return
}

func (r *UpdateRequest) checkParam(pageID uint64) error {
	if r.ID == 0 {
		return errno.New(errno.ErrValidation, nil).Add("need id.")
	}

	if r.ID != pageID {
		return errno.New(errno.ErrValidation, nil).Add("error id.")
	}

	if r.Status == "" {
		return errno.New(errno.ErrValidation, nil).Add("need status.")
	}

	if r.Status != "publish" && r.Status != "draft" && r.Status != "deleted" && r.Status != "restore" {
		return errno.New(errno.ErrValidation, nil).Add("error status.")
	}

	if r.Status == "publish" || r.Status == "draft" {
		if isExist := model.PageCheckSlugExist(r.ID, r.Slug); isExist == true {
			return errno.New(errno.ErrSlugExist, nil)
		}
	}

	return nil
}
