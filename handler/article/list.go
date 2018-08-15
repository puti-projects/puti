package article

import (
	"math"
	Response "puti/handler"
	"puti/pkg/constvar"
	"puti/pkg/errno"
	"puti/service"

	"github.com/gin-gonic/gin"
)

// List shows the article list in page
func List(c *gin.Context) {
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		Response.SendResponse(c, errno.ErrBind, nil)
		return
	}

	if r.Number == 0 {
		r.Number = constvar.DefaultLimit
	}

	infos, count, err := service.ListArticle(r.Title, r.Page, r.Number)
	if err != nil {
		Response.SendResponse(c, err, nil)
		return
	}

	number := uint64(r.Number)
	totalPage := math.Ceil(float64(count / number))

	Response.SendResponse(c, nil, ListResponse{
		TotalCount:  count,
		TotalPage:   uint64(totalPage),
		ArticleList: infos,
	})
}
