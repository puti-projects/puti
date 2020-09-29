package page

import (
	"strconv"

	"github.com/puti-projects/puti/internal/admin/api"
	"github.com/puti-projects/puti/internal/admin/service"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Get get page handler (page info detail)
func Get(c *gin.Context) {
	paegID, _ := strconv.Atoi(c.Param("id"))

	page, err := service.GetPageDetail(uint64(paegID))
	if err != nil {
		api.SendResponse(c, errno.ErrPageNotFount, nil)
		return
	}

	api.SendResponse(c, nil, page)
}
