package model

import "html/template"

// ShowArticle output article model
type ShowArticle struct {
	ID           uint64          `json:"id"`
	Title        template.HTML   `json:"title"`
	IfTop        bool            `json:"ifTop,omitempty"`
	Abstract     string          `json:"abstract"`
	GUID         string          `json:"url"`
	CoverPicture string          `json:"coverPictureUrl"`
	CommentCount uint64          `json:"commentCount"`
	ViewCount    uint64          `json:"viewCount"`
	PostedTime   string          `json:"postedTime"`
	Tags         []*ShowTag      `json:"tags"`
	Categories   []*ShowCategory `json:"categories"`
}

// ShowTag output tag model
type ShowTag struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

// ShowCategory output category model
type ShowCategory struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

// ShowWidgetLatestArticles latest article list for widget
type ShowWidgetLatestArticles struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	GUID         string `json:"GUID"`
	CommentCount string `json:"comment_count"`
	ViewCount    string `json:"view_count"`
	PostedTime   string `json:"posted_time"`
}

// ShowWidgetCategoryTreeNode category tree node for widget
type ShowWidgetCategoryTreeNode struct {
	TermID   uint64                        `json:"term_id"`
	Name     string                        `json:"name"`
	Slug     string                        `json:"slug"`
	Count    uint64                        `json:"count"`
	URL      string                        `json:"url"`
	Children []*ShowWidgetCategoryTreeNode `json:"children"`
}
