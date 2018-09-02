package media

import (
	Response "puti/handler"
	"puti/pkg/errno"
	"puti/service"

	"github.com/gin-gonic/gin"
)

// ListRequest is the media list request struct
type ListRequest struct {
	Limit int `form:"limit"`
	Page  int `form:"page"`
}

// ListResponse returns total number of media and current page of media
type ListResponse struct {
	TotalCount uint64               `json:"totalCount"`
	MediaList  []*service.MediaInfo `json:"mediaList"`
}

// List returns current page media list and the total number of media
func List(c *gin.Context) {
	var r ListRequest
	if err := c.ShouldBind(&r); err != nil {
		Response.SendResponse(c, errno.ErrBind, nil)
		return
	}

	infos, count, err := service.ListMedia(r.Limit, r.Page)
	if err != nil {
		Response.SendResponse(c, err, nil)
		return
	}

	Response.SendResponse(c, nil, ListResponse{
		TotalCount: count,
		MediaList:  infos,
	})
}
