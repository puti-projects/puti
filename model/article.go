package model

import (
	"time"
)

// ArticleModel is the struct model for article
type ArticleModel struct {
	Model

	UserID          uint64    `gorm:"column:user_id;not null"`
	PostType        string    `gorm:"column:post_type;not null"`
	Title           string    `gorm:"column:title;not null"`
	ContentMarkdown string    `gorm:"column:content_markdown;not null"`
	ContetnHTML     string    `gorm:"column:content_html;not null"`
	Slug            string    `gorm:"column:slug;not null"`
	ParentID        uint64    `gorm:"column:parent_id;not null"`
	Status          string    `gorm:"column:status;not null"`
	CommentStatus   uint64    `gorm:"column:comment_status;not null"`
	GUID            string    `gorm:"column:guid;not null"`
	CoverPicture    string    `gorm:"column:cover_picture;not null"`
	CommentCount    uint64    `gorm:"column:comment_count;not null"`
	ViewCount       uint64    `gorm:"column:view_count;not null"`
	PostDate        time.Time `gorm:"column:posted_time;not null"`
}

// TableName is the article table name in db
func (c *ArticleModel) TableName() string {
	return "pt_post"
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
