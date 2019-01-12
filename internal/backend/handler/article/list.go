package article

import (
	"math"

	Response "github.com/puti-projects/puti/internal/backend/handler"
	"github.com/puti-projects/puti/internal/backend/service"
	"github.com/puti-projects/puti/internal/pkg/constvar"
	"github.com/puti-projects/puti/internal/pkg/errno"

	"github.com/gin-gonic/gin"
)

// ListRequest is the article list request struct
type ListRequest struct {
	Title  string `form:"title"`
	Page   int    `form:"page"`
	Number int    `form:"number"`
	Sort   string `form:"sort"`
	Status string `form:"status"`
}

// ListResponse is the article list response struct
type ListResponse struct {
	TotalCount  uint64              `json:"totalCount"`
	TotalPage   uint64              `json:"totalPage"`
	ArticleList []*service.PostInfo `json:"articleList"`
}

// List return the article list in page
func List(c *gin.Context) {
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		Response.SendResponse(c, errno.ErrBind, nil)
		return
	}

	if r.Number == 0 {
		r.Number = constvar.DefaultLimit
	}

	infos, count, err := service.ListPost("article", r.Title, r.Page, r.Number, r.Sort, r.Status)
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
