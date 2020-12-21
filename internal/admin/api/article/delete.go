package article

import (
	"strconv"

	"github.com/puti-projects/puti/internal/admin/api"
	"github.com/puti-projects/puti/internal/admin/service"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Delete delete article handler
func Delete(c *gin.Context) {
	articleID, _ := strconv.Atoi(c.Param("id"))

	svc := service.New(c.Request.Context())
	if err := svc.DeletePost("article", uint64(articleID)); err != nil {
		api.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	api.SendResponse(c, nil, nil)
}
