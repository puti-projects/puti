package model

import (
	"github.com/puti-projects/puti/internal/common/utils"

	"github.com/jinzhu/gorm"
)

// MediaModel the definition of media model
type MediaModel struct {
	Model

	UserID      uint64 `gorm:"column:upload_user_id;not null"`
	Title       string `gorm:"column:title;not null"`
	Slug        string `gorm:"column:slug;not null"`
	Description string `gorm:"column:description;not null"`
	GUID        string `gorm:"column:guid;not null"`
	Type        string `gorm:"column:type;not null"`
	MimeType    string `gorm:"column:mime_type;not null"`
	Usage       string `gorm:"column:usage;not null"`
	Status      uint64 `gorm:"column:status;default:1;not null"`
}

// ResourceTypePicture resource type of picture
const ResourceTypePicture = "picture"

// StatusNormal normal status value
const StatusNormal = 1

// UsageDefault default usage
const UsageDefault = "common"

// UsageCover usage for cover
const UsageCover = "cover"

// TableName is the resource table name in db
func (c *MediaModel) TableName() string {
	return "pt_resource"
}

// BeforeCreate set values before create
// Set file type by mime-type
func (c *MediaModel) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("Type", utils.GetFileType(c.MimeType))
	return nil
}

// Create save the new media file info
func (c *MediaModel) Create() error {
	return DB.Local.Create(&c).Error
}

// Update update media info
func (c *MediaModel) Update() (err error) {
	if err = DB.Local.Model(&MediaModel{}).Save(c).Error; err != nil {
		return err
	}

	return nil
}

// DeleteMedia deletes the media info by id (not file)
func DeleteMedia(id uint64) error {
	media := MediaModel{}
	media.ID = id
	return DB.Local.Delete(&media).Error
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

// GetMediaByID get media info by id
func GetMediaByID(id uint64) (*MediaModel, error) {
	m := &MediaModel{}
	d := DB.Local.Where("status = 1 AND deleted_time is null AND id = ?", id).First(&m)
	return m, d.Error
}
