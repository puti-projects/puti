package media

import (
	"strconv"

	"github.com/puti-projects/puti/internal/admin/api"
	"github.com/puti-projects/puti/internal/admin/service"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Detail get media info detail handler
func Detail(c *gin.Context) {
	ID := c.Param("id")

	// Get media info by id
	svc := service.New(c.Request.Context())
	mediaID, _ := strconv.Atoi(ID)
	media, err := svc.GetMediaDetail(uint64(mediaID))
	if err != nil {
		api.SendResponse(c, errno.ErrMediaNotFound, nil)
		return
	}

	api.SendResponse(c, nil, media)
}
