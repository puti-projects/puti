package media

import (
	"strconv"

	"github.com/puti-projects/puti/internal/backend/api"
	"github.com/puti-projects/puti/internal/backend/service"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Detail get media info detail handler
func Detail(c *gin.Context) {
	ID := c.Param("id")

	// Get media info by id
	mediaID, _ := strconv.Atoi(ID)
	media, err := service.GetMediaDetail(uint64(mediaID))
	if err != nil {
		api.SendResponse(c, errno.ErrMediaNotFound, nil)
		return
	}

	api.SendResponse(c, nil, media)
}
