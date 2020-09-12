package model

import (
	"gorm.io/gorm"
)

// SubjectRelationshipsModel `pt_subject_relationships` 's struct
type SubjectRelationships struct {
	ObjectID  uint64 `gorm:"column:object_id;not null;primaryKey"`
	SubjectID uint64 `gorm:"column:subject_id;not null;primaryKey"`
	OrderNum  string `gorm:"column:order_num;not null"`
}

// TableName SubjectRelationshipsModel's binding db name
func (s *SubjectRelationships) TableName() string {
	return "pt_subject_relationships"
}

// GetAllByObjectID get article's related subject
func (s *SubjectRelationships) GetAllByObjectID(db *gorm.DB, objectID uint64) ([]*SubjectRelationships, error) {
	var subjectRelationships []*SubjectRelationships
	err := db.Where("`object_id` = ?", objectID).Find(&subjectRelationships).Error
	return subjectRelationships, err
}

// DeleteByCondition delete records by condition
func (s *SubjectRelationships) DeleteByCondition(db *gorm.DB, where string, whereArgs []interface{}) error {
	return db.Where(where, whereArgs...).Delete(s).Error
}
