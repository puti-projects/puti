package page

import (
	"strconv"

	"github.com/puti-projects/puti/internal/admin/api"
	"github.com/puti-projects/puti/internal/admin/service"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Delete delete page handler
func Delete(c *gin.Context) {
	pageID, _ := strconv.Atoi(c.Param("id"))

	if err := service.DeletePost("page", uint64(pageID)); err != nil {
		api.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	api.SendResponse(c, nil, nil)
}
