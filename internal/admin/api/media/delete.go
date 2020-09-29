package media

import (
	"strconv"

	"github.com/puti-projects/puti/internal/admin/api"
	"github.com/puti-projects/puti/internal/admin/service"

	"github.com/gin-gonic/gin"
)

// Delete delete the media info with deleted_at field (not file delete)
func Delete(c *gin.Context) {
	mediaID, _ := strconv.Atoi(c.Param("id"))

	if err := service.DeleteMedia(uint64(mediaID)); err != nil {
		api.SendResponse(c, err, nil)
		return
	}

	api.SendResponse(c, nil, nil)
}
