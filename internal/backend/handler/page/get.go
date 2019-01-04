package page

import (
	Response "github.com/puti-projects/puti/internal/backend/handler"
	"github.com/puti-projects/puti/internal/pkg/errno"
	"github.com/puti-projects/puti/internal/backend/service"

	"github.com/gin-gonic/gin"
)

// Get get article info detail
func Get(c *gin.Context) {
	paegID := c.Param("id")

	page, err := service.GetPageDetail(paegID)
	if err != nil {
		Response.SendResponse(c, errno.ErrPageNotFount, nil)
		return
	}

	Response.SendResponse(c, nil, page)
}
