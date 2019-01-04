package model

import (
	"time"
)

// PostModel is the struct model for post
type PostModel struct {
	Model

	UserID          uint64    `gorm:"column:user_id;not null"`
	PostType        string    `gorm:"column:post_type;not null"`
	Title           string    `gorm:"column:title;not null"`
	ContentMarkdown string    `gorm:"column:content_markdown;not null"`
	ContetnHTML     string    `gorm:"column:content_html;not null"`
	Slug            string    `gorm:"column:slug;not null"`
	ParentID        uint64    `gorm:"column:parent_id;not null"` // set to 0 now, use for draft history feature in the future
	Status          string    `gorm:"column:status;not null"`
	CommentStatus   uint64    `gorm:"column:comment_status;not null"`
	IfTop           uint64    `gorm:"column:if_top;not null"`
	GUID            string    `gorm:"column:guid;not null"`
	CoverPicture    string    `gorm:"column:cover_picture;not null"`
	CommentCount    uint64    `gorm:"column:comment_count;not null"`
	ViewCount       uint64    `gorm:"column:view_count;not null"`
	PostDate        time.Time `gorm:"column:posted_time;not null"`
}

// PostMetaModel meta data for post
type PostMetaModel struct {
	ID        uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id"`
	PostID    uint64 `gorm:"column:post_id;not null"`
	MetaKey   string `gorm:"column:meta_key;not null"`
	MetaValue string `gorm:"column:meta_value;not null"`
}

// TableName is the article table name in db
func (c *PostModel) TableName() string {
	return "pt_post"
}

// TableName is the article meta data table name in db
func (c *PostMetaModel) TableName() string {
	return "pt_post_meta"
}

// GetPost gets the post by post id
func GetPost(postID uint64) (*PostModel, error) {
	a := &PostModel{}
	d := DB.Local.Where("id = ? AND deleted_time is null", postID).First(&a)
	return a, d.Error
}

// GetPostMetaData gets the extral data of post
func GetPostMetaData(postID uint64) ([]*PostMetaModel, error) {
	am := []*PostMetaModel{}
	d := DB.Local.Where("post_id = ?", postID).Find(&am)
	return am, d.Error
}

// GetOnePostMetaData get one specific meta by metakey and post id
func GetOnePostMetaData(postID uint64, metaKey string) (*PostMetaModel, error) {
	am := &PostMetaModel{}
	d := DB.Local.Where("post_id = ? AND meta_key = ?", postID, metaKey).First(&am)
	return am, d.Error
}

// ListPost returns the posts list in condition
func ListPost(postType, title string, page, number int, sort, status string) ([]*PostModel, uint64, error) {
	posts := make([]*PostModel, 0)
	var count uint64

	where := "post_type = ? AND parent_id = ?"
	whereArgs := []interface{}{postType, 0}
	if "" != title {
		where += " AND title LIKE ?"
		whereArgs = append(whereArgs, "%"+title+"%")
	}

	if status != "" {
		where += " AND status= ?"
		whereArgs = append(whereArgs, status)
	}

	if err := DB.Local.Model(&PostModel{}).Where(where, whereArgs...).Count(&count).Error; err != nil {
		return posts, count, err
	}

	offset := (page - 1) * number
	var order string
	if sort != "" {
		order = "id " + sort
	} else {
		order = "id DESC"
	}

	if err := DB.Local.Where(where, whereArgs...).Offset(offset).Limit(number).Order(order).Find(&posts).Error; err != nil {
		return posts, count, err
	}

	return posts, count, nil
}

// PageCheckSlugExist check the slug if already exist
func PageCheckSlugExist(pageID uint64, Slug string) bool {
	post := &PostModel{}

	var ifNotFound bool
	if pageID > 0 {
		ifNotFound = DB.Local.Where("id != ? AND slug = ?", pageID, Slug).First(&post).RecordNotFound()
	} else {
		ifNotFound = DB.Local.Where("slug = ?", Slug).First(&post).RecordNotFound()
	}

	if ifNotFound {
		return false
	}

	return true
}
