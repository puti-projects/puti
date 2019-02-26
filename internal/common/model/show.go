package model

import "html/template"

// ShowArticle output article model for list
type ShowArticle struct {
	ID           uint64          `json:"id"`
	Title        string          `json:"title"`
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

// ShowLastAndNextArticle last or next article url model
type ShowLastOrNextArticle struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

// ShowArticleDetail article detail output model
type ShowArticleDetail struct {
	ID            uint64                 `json:"id"`
	Title         string                 `json:"title"`
	ContentHTML   template.HTML          `json:"content_html"`
	CommentStatus uint64                 `json:"comment_status"`
	GUID          string                 `json:"guid"`
	CommentCount  uint64                 `json:"commentCount"`
	ViewCount     uint64                 `json:"viewCount"`
	PostedTime    string                 `json:"posted_time"`
	MetaData      map[string]interface{} `json:"meta_date"`
	Tags          []*ShowTag             `json:"tags"`
	Categories    []*ShowCategory        `json:"categories"`
}

// ShowPageDetail page detail output model
type ShowPageDetail struct {
	ID            uint64                 `json:"id"`
	Title         string                 `json:"title"`
	ContentHTML   template.HTML          `json:"content_html"`
	CommentStatus uint64                 `json:"comment_status"`
	GUID          string                 `json:"guid"`
	CommentCount  uint64                 `json:"commentCount"`
	ViewCount     uint64                 `json:"viewCount"`
	PostedTime    string                 `json:"posted_time"`
	MetaData      map[string]interface{} `json:"meta_date"`
}

type ShowArchive struct {
	ID           uint64 `json:"id"`
	Title        string `json:"title"`
	GUID         string `json:"guid"`
	CommentCount uint64 `json:"commentCount"`
	ViewCount    uint64 `json:"viewCount"`
	PostedTime   string `json:"posted_time"`
	PostedDay    string `json:"posted_day"`
}
