package article

import "puti/model"

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
	TotalCount  uint64               `json:"totalCount"`
	TotalPage   uint64               `json:"totalPage"`
	ArticleList []*model.ArticleInfo `json:"articleList"`
}
