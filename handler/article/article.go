package article

import "puti/model"

// ListRequest is the article list request struct
type ListRequest struct {
	Title  string `form:"title"`
	Page   int    `form:"page"`
	Number int    `form:"number"`
}

// ListResponse is the article list response struct
type ListResponse struct {
	TotalCount  uint64               `json:"totalAccout"`
	TotalPage   uint64               `json:"totalPage"`
	ArticleList []*model.ArticleInfo `json:"articleList"`
}
