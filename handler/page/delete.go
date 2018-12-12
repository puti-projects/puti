package page

import (
	"strconv"

	Response "puti/handler"
	"puti/pkg/errno"
	"puti/service"

	"github.com/gin-gonic/gin"
)

// Delete delete page
func Delete(c *gin.Context) {
	pageID, _ := strconv.Atoi(c.Param("id"))

	if err := service.DeletePost("page", uint64(pageID)); err != nil {
		Response.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	Response.SendResponse(c, nil, nil)
}
