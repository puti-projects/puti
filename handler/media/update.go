package media

import (
	"strconv"

	Response "puti/handler"
	"puti/model"
	"puti/pkg/errno"
	"puti/service"
	"puti/util"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// UpdateRequest is the update media request params struct
type UpdateRequest struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
}

// Update update media info
func Update(c *gin.Context) {
	log.Info("Update function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})

	// Get user id
	userID, _ := strconv.Atoi(c.Param("id"))

	var r UpdateRequest

	if err := c.ShouldBind(&r); err != nil {
		Response.SendResponse(c, errno.ErrBind, nil)
		return
	}

	if r.Title == "" {
		Response.SendResponse(c, errno.ErrTitleEmpty, nil)
		return
	}

	r.ID = uint64(userID)

	media := &model.MediaModel{
		Model: model.Model{ID: r.ID},

		Title:       r.Title,
		Slug:        r.Slug,
		Description: r.Description,
	}

	// Update changed fields.
	if err := service.UpdateMedia(media); err != nil {
		Response.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	Response.SendResponse(c, nil, nil)
	return
}
