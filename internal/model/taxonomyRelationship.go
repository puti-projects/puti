package model

import (
	"gorm.io/gorm"
)

// TermRelationships `pt_term_relationships` 's struct taxomony raltionships with object
type TermRelationships struct {
	ObjectID       uint64 `gorm:"column:object_id;not null;primaryKey"`
	TermTaxonomyID uint64 `gorm:"column:term_taxonomy_id;not null;primaryKey"`
	TermOrder      string `gorm:"column:term_order;not null"`
}

// TableName TermRelationshipsModel's binding db name
func (c *TermRelationships) TableName() string {
	return "pt_term_relationships"
}

// Delete delete a record
func (t *TermRelationships) Delete(db *gorm.DB) error {
	return db.Delete(t).Error
}

// DeleteByCondition delete records by condition
func (t *TermRelationships) DeleteByCondition(db *gorm.DB, where string, whereArgs []interface{}) error {
	return db.Where(where, whereArgs...).Delete(t).Error
}
