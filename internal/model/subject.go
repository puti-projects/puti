package model

import (
	"database/sql"

	"gorm.io/gorm"
)

// SubjectModel the definition of subject model
type Subject struct {
	Model

	ParentID    uint64       `gorm:"column:parent_id;not null"`
	Name        string       `gorm:"column:name;not null"`
	Slug        string       `gorm:"column:slug;not null"`
	Description string       `gorm:"column:description;not null"`
	CoverImage  uint64       `gorm:"column:cover_image;not null"`
	IsEnd       uint64       `gorm:"column:is_end;not null"`
	Count       uint64       `gorm:"column:count;not null"`
	LastUpdated sql.NullTime `gorm:"column:last_updated"`
}

// TableName is the resource table name in db
func (s *Subject) TableName() string {
	return "pt_subject"
}

// Create creates a new subject
func (s *Subject) Create(db *gorm.DB) error {
	return db.Create(s).Error
}

// Update update subject
func (s *Subject) Update(db *gorm.DB) error {
	return db.Save(s).Error
}

// Save update subject
func (s *Subject) Save(db *gorm.DB) error {
	return db.Save(s).Error
}

// Delete delete subject
func (s *Subject) Delete(db *gorm.DB) error {
	return db.Delete(s).Error
}

// GetByID get subject by ID
func (s *Subject) GetByID(db *gorm.DB) error {
	return db.First(s, s.ID).Error
}

// GetAll get all subjects
func (s *Subject) GetAll(db *gorm.DB) ([]*Subject, error) {
	var subjects []*Subject
	err := db.Find(&subjects).Error
	return subjects, err
}

// CheckSubjectNameExist check the subject name if is already exist
func (s *Subject) CheckSubjectNameExist(db *gorm.DB) bool {
	var count int64 = 0
	if s.ID > 0 {
		db.Model(s).Where("`id` != ? AND `name` = ?", s.ID, s.Name).Count(&count)
	} else {
		db.Model(s).Where("`name` = ?", s.Name).Count(&count)
	}

	if count > 0 {
		return true
	}

	return false
}

// IfSubjectHasChild check if subject has children
func (s *Subject) IfSubjectHasChild(db *gorm.DB, subjectID uint64) bool {
	var count int64
	db.Model(&s).Where("`parent_id` = ?", subjectID).Count(&count)

	if count > 0 {
		return true
	}

	return false
}
