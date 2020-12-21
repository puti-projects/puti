package media

import (
	"github.com/puti-projects/puti/internal/admin/api"
	"github.com/puti-projects/puti/internal/admin/service"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// List media list handler
func List(c *gin.Context) {
	var r service.MediaListRequest
	if err := c.ShouldBind(&r); err != nil {
		api.SendResponse(c, errno.ErrBind, nil)
		return
	}

	svc := service.New(c.Request.Context())
	infos, count, err := svc.ListMedia(r.Limit, r.Page)
	if err != nil {
		api.SendResponse(c, err, nil)
		return
	}

	api.SendResponse(c, nil, service.MediaListResponse{
		TotalCount: count,
		MediaList:  infos,
	})
}
