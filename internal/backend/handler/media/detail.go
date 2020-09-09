package media

import (
	"strconv"

	Response "github.com/puti-projects/puti/internal/backend/handler"
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
		Response.SendResponse(c, errno.ErrMediaNotFound, nil)
		return
	}

	Response.SendResponse(c, nil, media)
}
