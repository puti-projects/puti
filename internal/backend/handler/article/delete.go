package article

import (
	"strconv"

	Response "github.com/puti-projects/puti/internal/backend/handler"
	"github.com/puti-projects/puti/internal/pkg/errno"
	"github.com/puti-projects/puti/internal/backend/service"

	"github.com/gin-gonic/gin"
)

// Delete delete article and it's relationship
func Delete(c *gin.Context) {
	articleID, _ := strconv.Atoi(c.Param("id"))

	if err := service.DeletePost("article", uint64(articleID)); err != nil {
		Response.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	Response.SendResponse(c, nil, nil)
}
