package model

import (
	"database/sql"
	"gorm.io/gorm"
)

// Knowledge the definition of knowledge base model
// Note: because of the cover image will be uploaded first by a api, here do not consider to use a association mode
type Knowledge struct {
	Model

	Name        string       `gorm:"column:name;not null"`
	Slug        string       `gorm:"column:slug;not null"`
	Type        string       `gorm:"column:type;not null"`
	Description string       `gorm:"column:description;not null"`
	CoverImage  uint64       `gorm:"column:cover_image;not null;default:0"`
	Status      uint8        `gorm:"column:status;not null;default:1"`
	LastUpdated sql.NullTime `gorm:"column:last_updated;default:null"`
}

const (
	// KnowledgeTypeNote knowledge type of note
	KnowledgeTypeNote = "note"
	// KnowledgeTypeDoc knowledge type of document
	KnowledgeTypeDoc = "doc"
)

// TableName is the Knowledge table name in db
func (k *Knowledge) TableName() string {
	return "pt_knowledge"
}

// Create creates a new knowledge base
func (k *Knowledge) Create(db *gorm.DB) error {
	return db.Create(k).Error
}

// Save save knowledge base info
func (k *Knowledge) Save(db *gorm.DB) error {
	return db.Save(k).Error
}

// GetByID get knowledge base by ID
func (k *Knowledge) GetByID(db *gorm.DB) error {
	return db.First(k, k.ID).Error
}

// GetAll get all knowledge bases
//func (k *Knowledge) GetAll(db *gorm.DB) ([]*Knowledge, error) {
//	var ks []*Knowledge
//	err := db.Find(&ks).Error
//	return ks, err
//}
