package page

import (
	"math"

	"github.com/puti-projects/puti/internal/backend/api"
	"github.com/puti-projects/puti/internal/backend/service"
	"github.com/puti-projects/puti/internal/pkg/constvar"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// List return the page list in page
func List(c *gin.Context) {
	var r service.PageListRequest
	if err := c.Bind(&r); err != nil {
		api.SendResponse(c, errno.ErrBind, nil)
		return
	}

	if r.Number == 0 {
		r.Number = constvar.DefaultLimit
	}

	infos, count, err := service.ListPage("page", &r)
	if err != nil {
		api.SendResponse(c, err, nil)
		return
	}

	number := int64(r.Number)
	totalPage := math.Ceil(float64(count / number))

	api.SendResponse(c, nil, service.PageListResponse{
		TotalCount: count,
		TotalPage:  uint64(totalPage),
		PageList:   infos,
	})
}
