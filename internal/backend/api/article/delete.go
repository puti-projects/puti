package article

import (
	"strconv"

	"github.com/puti-projects/puti/internal/backend/api"
	"github.com/puti-projects/puti/internal/backend/service"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Delete delete article handler
func Delete(c *gin.Context) {
	articleID, _ := strconv.Atoi(c.Param("id"))

	if err := service.DeletePost("article", uint64(articleID)); err != nil {
		api.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	api.SendResponse(c, nil, nil)
}
