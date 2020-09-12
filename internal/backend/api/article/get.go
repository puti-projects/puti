package article

import (
	"strconv"

	"github.com/puti-projects/puti/internal/backend/api"
	"github.com/puti-projects/puti/internal/backend/service"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Get get article handler (article info in detail)
func Get(c *gin.Context) {
	articleID, _ := strconv.Atoi(c.Param("id"))

	article, err := service.GetArticleDetail(uint64(articleID))
	if err != nil {
		api.SendResponse(c, errno.ErrArticleNotFount, nil)
		return
	}

	api.SendResponse(c, nil, article)
}
