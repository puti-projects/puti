package model

import (
	"sync"
	"time"
)

// ArticleInfo is article info for article list
type ArticleInfo struct {
	ID           uint64 `json:"id"`
	UserID       uint64 `json:"userId"`
	Title        string `json:"title"`
	Status       string `json:"status"`
	PostDate     string `json:"post_date"`
	CommentCount uint64 `json:"comment_count"`
	ViewCount    uint64 `json:"view_count"`
}

// ArticleList article list
type ArticleList struct {
	Lock  *sync.Mutex
	IDMap map[uint64]*ArticleInfo
}

// ArticleModel is the struct model for article
type ArticleModel struct {
	ID              uint64    `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	UserID          uint64    `gorm:"column:user_id;not null" json:"user_id"`
	PostType        string    `gorm:"column:post_type;not null" json:"post_type"`
	Title           string    `gorm:"column:title;not null" json:"title"`
	ContentMarkdown string    `gorm:"column:content_markdown;not null" json:"content_markdown"`
	ContetnHTML     string    `gorm:"column:content_html;not null" json:"content_html"`
	Slug            string    `gorm:"column:slug;not null" json:"slug"`
	ParentID        uint64    `gorm:"column:parent_id;not null" json:"parent_id"`
	Status          string    `gorm:"column:status;not null" json:"status"`
	CommentStatus   uint64    `gorm:"column:comment_status;not null" json:"comment_status"`
	PostDate        time.Time `gorm:"column:post_date;not null" json:"post_date"`
	PostDateGmt     time.Time `gorm:"column:post_date_gmt;not null" json:"post_date_gmt"`
	PostModified    time.Time `gorm:"column:post_modified;not null" json:"post_modified"`
	PostModofiedGmt time.Time `gorm:"column:post_modified_gmt;not null" json:"post_modified_gmt"`
	GUID            string    `gorm:"column:guid;not null" json:"guid"`
	CoverPicture    string    `gorm:"column:cover_picture;not null" json:"cover_picture"`
	CommentCount    uint64    `gorm:"column:comment_count;not null" json:"comment_count"`
	ViewCount       uint64    `gorm:"column:view_count;not null" json:"view_count"`
}

// TableName is the article table name in db
func (c *ArticleModel) TableName() string {
	return "pt_posts"
}

// ListArticle shows the articles in condition
func ListArticle(title string, page, number int, sort, status string) ([]*ArticleModel, uint64, error) {
	articles := make([]*ArticleModel, 0)
	var count uint64

	where := "post_type = ? AND parent_id = ?"
	whereArgs := []interface{}{"post", 0}
	if "" != title {
		where += " AND title LIKE ?"
		whereArgs = append(whereArgs, "%"+title+"%")
	}

	if status != "" {
		where += " AND status= ?"
		whereArgs = append(whereArgs, status)
	}

	if err := DB.Local.Model(&ArticleModel{}).Where(where, whereArgs...).Count(&count).Error; err != nil {
		return articles, count, err
	}

	offset := (page - 1) * number
	var order string
	if sort != "" {
		order = "id " + sort
	} else {
		order = "id DESC"
	}

	if err := DB.Local.Where(where, whereArgs...).Offset(offset).Limit(number).Order(order).Find(&articles).Error; err != nil {
		return articles, count, err
	}

	return articles, count, nil
}
