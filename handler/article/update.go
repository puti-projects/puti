package article

import (
	"strconv"

	Response "puti/handler"
	"puti/model"
	"puti/pkg/errno"
	"puti/service"
	"puti/utils"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

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
}

func Update(c *gin.Context) {
	log.Info("Update function called.", lager.Data{"X-Request-Id": utils.GetReqID(c)})

	// Get term id
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

	article := &model.ArticleModel{
		Model:           model.Model{ID: r.ID},
		Title:           r.Title,
		ContentMarkdown: r.Content,
		ContetnHTML:     r.ContentHTML,
		Status:          r.Status,
		CommentStatus:   r.CommentStatus,
		IfTop:           r.IfTop,
		CoverPicture:    r.CoverPicture,
		PostDate:        utils.StringToTime("2006-01-02 15:04:05", r.PostedTime),
	}

	// Update changed fields.
	if err := service.UpdateArticle(article, r.Description, r.Category, r.Tag); err != nil {
		Response.SendResponse(c, errno.ErrDatabase, nil)
		return
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

	return nil
}
