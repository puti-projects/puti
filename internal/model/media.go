package model

import (
	"github.com/puti-projects/puti/internal/utils"

	"gorm.io/gorm"
)

// MediaModel the definition of media model
type Media struct {
	Model

	UserID      uint64 `gorm:"column:upload_user_id;not null"`
	Title       string `gorm:"column:title;not null"`
	Slug        string `gorm:"column:slug;not null"`
	Description string `gorm:"column:description;not null"`
	GUID        string `gorm:"column:guid;not null"`
	Type        string `gorm:"column:type;not null;default:picture"`
	MimeType    string `gorm:"column:mime_type;not null"`
	Usage       string `gorm:"column:usage;not null"`
	Status      uint64 `gorm:"column:status;not null;default:1"`
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
func (m *Media) TableName() string {
	return "pt_resource"
}

// BeforeCreate set values before create
// Set file type by mime-type
func (m *Media) BeforeCreate(tx *gorm.DB) (err error) {
	m.Type = utils.GetFileType(m.MimeType)
	return
}

// Create save the new media file info
func (m *Media) Create(db *gorm.DB) error {
	return db.Create(m).Error
}

// Update update media info
func (m *Media) Update(db *gorm.DB) (err error) {
	return db.Save(m).Error
}

// GetByID get media info by ID
func (m *Media) GetByID(db *gorm.DB) error {
	return db.Where("`status` = 1 AND `deleted_time` is null AND `id` = ?", m.ID).First(m).Error
}

// Delete delete the media info by id (not file right now)
func (m *Media) Delete(db *gorm.DB) error {
	return db.Delete(m).Error
}

// Count count media in condition
func (m *Media) Count(db *gorm.DB, where string, whereArgs []interface{}) (count int64, err error) {
	if whereArgs != nil {
		err = db.Model(m).Where(where, whereArgs...).Count(&count).Error
		return
	}
	err = db.Model(m).Where(where).Count(&count).Error
	return
}

// List get media list
func (m *Media) List(db *gorm.DB, where string, whereArgs []interface{}, offset, limit int) (medias []*Media, err error) {
	medias = make([]*Media, 0)
	if whereArgs != nil {
		err = db.Where(where, whereArgs...).Offset(offset).Limit(limit).Order("created_time DESC").Find(&medias).Error
		return
	}

	err = db.Where(where).Offset(offset).Limit(limit).Order("created_time DESC").Find(&medias).Error
	return medias, err
}

// TotalNumber get total number of media
func (m *Media) TotalNumber(db *gorm.DB) (totalMedia int64, err error) {
	err = db.Model(m).
		Where("`deleted_time` is null").
		Count(&totalMedia).Error
	return
}
