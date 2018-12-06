package article

import (
	"strconv"

	Response "puti/handler"
	"puti/pkg/errno"
	"puti/service"

	"github.com/gin-gonic/gin"
)

// Delete delete article and it's relationship
func Delete(c *gin.Context) {
	articleID, _ := strconv.Atoi(c.Param("id"))

	if err := service.DeleteArticle(uint64(articleID)); err != nil {
		Response.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	Response.SendResponse(c, nil, nil)
}
