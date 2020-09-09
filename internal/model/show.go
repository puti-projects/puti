package model

import (
	"database/sql"
	"html/template"
)

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
// Use {{ formatNullTime .PostedTime "2006-01-02 15:04" }} to decode the time
type ShowWidgetLatestArticles struct {
	ID           string        `json:"id"`
	Title        string        `json:"title"`
	GUID         string        `json:"GUID"`
	CommentCount string        `json:"comment_count"`
	ViewCount    string        `json:"view_count"`
	PostedTime   *sql.NullTime `json:"posted_time"`
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

// ShowLastOrNextArticle last or next article url model
type ShowLastOrNextArticle struct {
	Title string `json:"title"`
	URL   string `json:"url"`
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

// ShowArchive archive item
type ShowArchive struct {
	ID           uint64 `json:"id"`
	Title        string `json:"title"`
	GUID         string `json:"guid"`
	CommentCount uint64 `json:"commentCount"`
	ViewCount    uint64 `json:"viewCount"`
	PostedTime   string `json:"posted_time"`
	PostedDay    string `json:"posted_day"`
}

// ShowSubjectInfo show subejcts info output model
type ShowSubjectInfo struct {
	ID            uint64 `json:"id"`
	ParentURL     string `json:"parent_url"`
	Name          string `json:"name"`
	Slug          string `json:"slug"`
	Description   string `json:"description"`
	CoverImageURL string `json:"cover_image_url"`
	Count         string `json:"count"`
}

// ShowSubjectList show subejcts list output model
type ShowSubjectList struct {
	ID                uint64 `json:"id"`
	URL               string `json:"url"`
	Name              string `json:"name"`
	Slug              string `json:"slug"`
	Description       string `json:"description"`
	CoverImageURL     string `json:"cover_image_url"`
	Count             uint64 `json:"count"`
	LastUpdated       string `json:"last_updated"`
	SubLastUpdatedDay string `json:"sub_last_updated"`
}
