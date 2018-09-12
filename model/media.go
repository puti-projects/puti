package model

import (
	"puti/util"

	"github.com/jinzhu/gorm"
)

// MediaModel the definition of media model
type MediaModel struct {
	Model

	UserID      uint64 `gorm:"column:upload_user_id;not null"`
	PostID      uint64 `gorm:"column:post_id;default:0;not null"`
	Title       string `gorm:"columntitle:;not null"`
	Slug        string `gorm:"column:slug;not null"`
	Description string `gorm:"column:description;not null"`
	GUID        string `gorm:"column:guid;not null"`
	Type        string `gorm:"column:type;not null"`
	MimeType    string `gorm:"column:mime_type;not null"`
	Status      uint64 `gorm:"column:status;default:1;not null"`
}

// TableName is the resource table name in db
func (c *MediaModel) TableName() string {
	return "pt_resources"
}

// BeforeCreate set values before create
// Set file type by mime-type
func (c *MediaModel) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("Type", util.GetFileType(c.MimeType))
	return nil
}

// Create save the new media file info
func (c *MediaModel) Create() error {
	return DB.Local.Create(&c).Error
}

// ListMedia returns the media list in condition
func ListMedia(limit, page int) ([]*MediaModel, uint64, error) {
	medias := make([]*MediaModel, 0)
	var count uint64

	where := "deleted_time is null"
	if err := DB.Local.Model(&MediaModel{}).Where(where).Count(&count).Error; err != nil {
		return medias, count, err
	}

	offset := (page - 1) * limit
	if err := DB.Local.Where(where).Offset(offset).Limit(limit).Order("created_time DESC").Find(&medias).Error; err != nil {
		return medias, count, err
	}

	return medias, count, nil
}
