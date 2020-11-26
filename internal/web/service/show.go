package service

import (
	"database/sql"
	"html/template"
)

// ShowArticle output article model for list
type ShowArticle struct {
	ID           uint64
	Title        string
	IfTop        bool
	Abstract     string
	GUID         string
	CoverPicture string
	CommentCount uint64
	ViewCount    uint64
	PostedTime   string
	Tags         []*ShowTag
	Categories   []*ShowCategory
}

// ShowTag output tag model
type ShowTag struct {
	Title string
	URL   string
}

// ShowCategory output category model
type ShowCategory struct {
	Title string
	URL   string
}

// ShowWidgetLatestArticles latest article list for widget
// Use {{ formatNullTime .PostedTime "2006-01-02 15:04" }} to decode the time
type ShowWidgetLatestArticles struct {
	ID           string
	Title        string
	GUID         string
	CommentCount string
	ViewCount    string
	PostedTime   sql.NullTime
}

// ShowWidgetCategoryTreeNode category tree node for widget
type ShowWidgetCategoryTreeNode struct {
	TermID   uint64
	Name     string
	Slug     string
	Count    uint64
	URL      string
	Children []*ShowWidgetCategoryTreeNode
}

// ShowLastOrNextArticle last or next article url model
type ShowLastOrNextArticle struct {
	Title string
	URL   string
}

// ShowArticleDetail article detail output model
type ShowArticleDetail struct {
	ID            uint64
	Title         string
	ContentHTML   template.HTML
	CommentStatus uint64
	GUID          string
	CommentCount  uint64
	ViewCount     uint64
	PostedTime    string
	MetaData      map[string]interface{}
	Tags          []*ShowTag
	Categories    []*ShowCategory
}

// ShowPageDetail page detail output model
type ShowPageDetail struct {
	ID            uint64
	Title         string
	ContentHTML   template.HTML
	CommentStatus uint64
	GUID          string
	CommentCount  uint64
	ViewCount     uint64
	PostedTime    string
	MetaData      map[string]interface{}
}

// ShowArchive archive item
type ShowArchive struct {
	ID           uint64
	Title        string
	GUID         string
	CommentCount uint64
	ViewCount    uint64
	PostedTime   string
	PostedDay    string
}

// ShowSubjectInfo show subjects info output model
type ShowSubjectInfo struct {
	ID            uint64
	ParentURL     string
	Name          string
	Slug          string
	Description   string
	CoverImageURL string
	Count         string
}

// ShowSubjectList show subjects list output model
type ShowSubjectList struct {
	ID                uint64
	URL               string
	Name              string
	Slug              string
	Description       string
	CoverImageURL     string
	Count             uint64
	LastUpdated       string
	SubLastUpdatedDay string
}

// ShowKnowledgeList show knowledge list output model
type ShowKnowledgeList struct {
	ID            uint64
	URL           string
	Name          string
	Slug          string
	Description   string
	CoverImageURL string
	UpdatedTime   string
}

// ShowKnowledgeInfo knowledge info output model
type ShowKnowledgeInfo struct {
	ID          uint64
	Name        string
	LastUpdated string
}

// ShowWidgetCategoryTreeNode category tree node for widget
type ShowKnowledgeItemTreeNode struct {
	ID       uint64
	Symbol   uint64
	Title    string
	URL      string
	Level    uint64
	Index    int64
	Children []*ShowKnowledgeItemTreeNode
}

// ShowKnowledgeItemContent knowledge item content
type ShowKnowledgeItemContent struct {
	Symbol  uint64
	Title   string
	Content template.HTML
}
