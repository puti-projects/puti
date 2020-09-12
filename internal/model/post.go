package model

import (
	"database/sql"
	"errors"

	"gorm.io/gorm"
)

// PostModel is the struct model for post
type Post struct {
	Model

	UserID          uint64       `gorm:"column:user_id;not null"`
	PostType        string       `gorm:"column:post_type;not null;default:article"`
	Title           string       `gorm:"column:title;not null"`
	ContentMarkdown string       `gorm:"column:content_markdown;not null"`
	ContentHTML     string       `gorm:"column:content_html;not null"`
	Slug            string       `gorm:"column:slug;not null"`
	ParentID        uint64       `gorm:"column:parent_id;not null"` // set to 0 now, use for draft history feature in the future
	Status          string       `gorm:"column:status;not null;default:publish"`
	CommentStatus   uint64       `gorm:"column:comment_status;not null;default:1"`
	IfTop           uint64       `gorm:"column:if_top;not null"`
	GUID            string       `gorm:"column:guid;not null"`
	CoverPicture    string       `gorm:"column:cover_picture;not null"`
	CommentCount    uint64       `gorm:"column:comment_count;not null"`
	ViewCount       uint64       `gorm:"column:view_count;not null"`
	PostDate        sql.NullTime `gorm:"column:posted_time;not null"`
}

// PostMetaModel meta data for post
type PostMeta struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;column:id"`
	PostID    uint64 `gorm:"column:post_id;not null"`
	MetaKey   string `gorm:"column:meta_key;not null"`
	MetaValue string `gorm:"column:meta_value;not null"`
}

const (
	// PostTypeArticle post type of article
	PostTypeArticle = "article"
	// PostTypePage post type of page
	PostTypePage = "page"
	// PostStatusPublish post status of published post
	PostStatusPublish = "publish"
	// PostStatusDraft post status of draft post
	PostStatusDraft = "draft"
	// PostStatusDeleted post status of deleted post
	PostStatusDeleted = "deleted"
)

// TableName is the article table name in db
func (p *Post) TableName() string {
	return "pt_post"
}

// TableName is the article meta data table name in db
func (p *PostMeta) TableName() string {
	return "pt_post_meta"
}

// Create creates a post
func (p *Post) Create(db *gorm.DB) error {
	return db.Create(p).Error
}

// Create
func (p *PostMeta) Create(db *gorm.DB) error {
	return db.Create(p).Error
}

// Save update a post
func (p *Post) Save(db *gorm.DB) error {
	return db.Save(p).Error
}

// Save update a post meata
func (p *PostMeta) Save(db *gorm.DB) error {
	return db.Save(p).Error
}

// Delete delete a post by ID
func (p *Post) Delete(db *gorm.DB) error {
	return db.Delete(p).Error
}

// GetByID get post by ID
func (p *Post) GetByID(db *gorm.DB) error {
	return db.Where("`deleted_time` is null").First(p, p.ID).Error
}

// GetAllByPostID get all meta data by post id
func (p *PostMeta) GetAllByPostID(db *gorm.DB) ([]*PostMeta, error) {
	pm := []*PostMeta{}
	err := db.Where("`post_id` = ?", p.PostID).Find(&pm).Error
	return pm, err
}

// GetOneByPostID get one specific meta by metakey and post id
func (p *PostMeta) GetOneByPostID(db *gorm.DB) error {
	if "" != p.MetaKey {
		return db.Where("`post_id` = ? AND `meta_key` = ?", p.PostID, p.MetaKey).First(&p).Error
	}
	return nil
}

// Count count user
func (p *Post) Count(db *gorm.DB, where string, whereArgs []interface{}) (int64, error) {
	var count int64
	err := db.Model(p).Where(where, whereArgs...).Count(&count).Error
	return count, err
}

// List get user list
func (p *Post) List(db *gorm.DB, where string, whereArgs []interface{}, offset, number int, order string) ([]*Post, error) {
	post := make([]*Post, 0)
	err := db.Where(where, whereArgs...).Offset(offset).Limit(number).Order(order).Find(&post).Error
	return post, err
}

// PageCheckSlugExist check the slug if already exist
// ErrRecordNotFound => False
// exist => True
func (p *Post) CheckSlug(db *gorm.DB) bool {
	var err error
	if p.ID > 0 {
		err = db.Where("`id` != ? AND `slug` = ?", p.ID, p.Slug).First(&p).Error
	} else {
		err = db.Where("`slug` = ?", p.Slug).First(&p).Error
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}

	return true
}

// TotalView gte total view of all post
func (p *Post) TotalView(db *gorm.DB) (totalViews int64, err error) {
	row := db.Model(p).
		Where("`status` != ? AND `deleted_time` is null", "deleted").
		Select("sum(`view_count`) as total_views").
		Row()
	err = row.Scan(&totalViews)
	return
}

// TotalNumber count total number of post by post type
func (p *Post) TotalNumber(db *gorm.DB, postType string) (totalPost int64, err error) {
	err = db.Model(p).
		Where("`post_type` = ? AND `status` != ? AND `deleted_time` is null", postType, "deleted").
		Count(&totalPost).Error
	return
}
