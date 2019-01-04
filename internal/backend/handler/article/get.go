package article

import (
	Response "github.com/puti-projects/puti/internal/backend/handler"
	"github.com/puti-projects/puti/internal/pkg/errno"
	"github.com/puti-projects/puti/internal/backend/service"

	"github.com/gin-gonic/gin"
)

// Get get article info in detail
func Get(c *gin.Context) {
	articleID := c.Param("id")

	article, err := service.GetArticleDetail(articleID)
	if err != nil {
		Response.SendResponse(c, errno.ErrArticleNotFount, nil)
		return
	}

	Response.SendResponse(c, nil, article)
}
