package media

import (
	"strconv"

	Response "github.com/puti-projects/puti/internal/backend/handler"
	"github.com/puti-projects/puti/internal/model"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Delete delete the media info with deleted_at field (not file delete)
func Delete(c *gin.Context) {
	mediaID, _ := strconv.Atoi(c.Param("id"))

	if err := model.DeleteMedia(uint64(mediaID)); err != nil {
		Response.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	Response.SendResponse(c, nil, nil)
}
