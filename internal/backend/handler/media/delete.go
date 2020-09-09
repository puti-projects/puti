package media

import (
	"strconv"

	Response "github.com/puti-projects/puti/internal/backend/handler"
	"github.com/puti-projects/puti/internal/backend/service"

	"github.com/gin-gonic/gin"
)

// Delete delete the media info with deleted_at field (not file delete)
func Delete(c *gin.Context) {
	mediaID, _ := strconv.Atoi(c.Param("id"))

	if err := service.DeleteMedia(uint64(mediaID)); err != nil {
		Response.SendResponse(c, err, nil)
		return
	}

	Response.SendResponse(c, nil, nil)
}
