package media

import (
	"strconv"

	Response "puti/handler"
	"puti/pkg/errno"
	"puti/service"

	"github.com/gin-gonic/gin"
)

// Detail get media info
func Detail(c *gin.Context) {
	ID := c.Param("id")

	// Get media info by id
	mediaID, err := strconv.Atoi(ID)
	media, err := service.GetMedia(uint64(mediaID))
	if err != nil {
		Response.SendResponse(c, errno.ErrMediaNotFound, nil)
		return
	}

	Response.SendResponse(c, nil, media)
}
