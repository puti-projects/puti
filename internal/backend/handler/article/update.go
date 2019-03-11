package article

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

// UpdateRequest struct for update article
type UpdateRequest struct {
	ID            uint64   `json:"id"`
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

// Update update article
// Delete and restore article are also in this function and it depends on the 'status'
func Update(c *gin.Context) {
	logger.Info("article update function called", zap.String("X-request-Id", utils.GetReqID(c)))

	// Get article id
	ID, _ := strconv.Atoi(c.Param("id"))

	var r UpdateRequest
	if err := c.ShouldBind(&r); err != nil {
		Response.SendResponse(c, errno.ErrBind, nil)
		return
	}

	articleID := uint64(ID)

	// check params
	if err := r.checkParam(articleID); err != nil {
		Response.SendResponse(c, err, nil)
		return
	}

	if r.Status == "deleted" {
		if err := service.TrashPost(articleID); err != nil {
			Response.SendResponse(c, errno.ErrDatabase, nil)
			return
		}
	} else if r.Status == "restore" {
		if err := service.RestorePost(articleID); err != nil {
			Response.SendResponse(c, errno.ErrDatabase, nil)
			return
		}
	} else {
		article := &model.PostModel{
			Model:           model.Model{ID: r.ID},
			Title:           r.Title,
			ContentMarkdown: r.Content,
			ContentHTML:     r.ContentHTML,
			Status:          r.Status,
			CommentStatus:   r.CommentStatus,
			IfTop:           r.IfTop,
			CoverPicture:    r.CoverPicture,
			PostDate:        utils.StringToNullTime("2006-01-02 15:04:05", r.PostedTime),
		}

		// Update changed fields.
		if err := service.UpdateArticle(article, r.Description, r.Category, r.Tag, r.Subject); err != nil {
			Response.SendResponse(c, errno.ErrDatabase, nil)
			return
		}
	}

	Response.SendResponse(c, nil, nil)
	return
}

func (r *UpdateRequest) checkParam(articleID uint64) error {
	if r.ID == 0 {
		return errno.New(errno.ErrValidation, nil).Add("need id.")
	}

	if r.ID != articleID {
		return errno.New(errno.ErrValidation, nil).Add("error id.")
	}

	if r.Status == "" {
		return errno.New(errno.ErrValidation, nil).Add("need status.")
	}

	if r.Status != "publish" && r.Status != "draft" && r.Status != "deleted" && r.Status != "restore" {
		return errno.New(errno.ErrValidation, nil).Add("error status.")
	}

	return nil
}
