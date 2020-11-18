package model

import (
	"gorm.io/gorm"
)

// TermTaxonomy `pt_term_taxonomy`'s struct
type TermTaxonomy struct {
	ID           uint64 `gorm:"column:term_taxonomy_id;not null;primaryKey"`
	Term         Term   `gorm:"foreignKey:TermID;references:ID"`
	TermID       uint64 `gorm:"column:term_id;not null"`
	ParentTermID uint64 `gorm:"column:parent_term_id;not null"`
	Level        uint64 `gorm:"column:level;not null;default:1"`
	Taxonomy     string `gorm:"column:taxonomy;not null"`
	TermGroup    uint64 `gorm:"column:term_group;default:0;not null"`
}

// Term `pt_terms`'s struct for taxomony info
type Term struct {
	ID          uint64 `gorm:"column:term_id;not null;primary_key"`
	Name        string `gorm:"column:name;not null"`
	Slug        string `gorm:"column:slug;not null"`
	Description string `gorm:"column:description;not null"`
	Count       uint64 `gorm:"column:count;not null"`
}

// DefaultUnCategorizedID the default category's ID
const DefaultUnCategorizedID = 1

// TableName TermTaxonomyModel's binding db name
func (t *TermTaxonomy) TableName() string {
	return "pt_term_taxonomy"
}

// TableName TermModel's binding db name
func (t *Term) TableName() string {
	return "pt_term"
}

// Create creates a new taxonomy
func (t *TermTaxonomy) Create(db *gorm.DB) error {
	return db.Create(t).Error
}

// Get get all by condition
func (t *TermTaxonomy) Get(db *gorm.DB, where string, whereArgs []interface{}) ([]*TermTaxonomy, error) {
	termTaxonomy := make([]*TermTaxonomy, 0)
	err := db.Where(where, whereArgs...).Find(&termTaxonomy).Error
	return termTaxonomy, err
}

// Count count taxonomy
func (t *TermTaxonomy) Count(db *gorm.DB, where string, whereArgs []interface{}) (int64, error) {
	var count int64
	err := db.Model(t).Where(where, whereArgs...).Count(&count).Error
	return count, err
}

// GetByTermID get term taxonomy by term id
func (t *TermTaxonomy) GetByTermID(db *gorm.DB) error {
	if t.Taxonomy != "" {
		return db.Where("`term_id` = ? AND `taxonomy` = ?", t.TermID, t.Taxonomy).First(t).Error
	}
	return db.Where("`term_id` = ?", t.TermID).First(t).Error
}

func (t *TermTaxonomy) GetColumnByTermID(db *gorm.DB, columns ...string) error {
	return db.Select(columns).Where("`term_id` = ?", t.TermID).First(t).Error
}

// GetByTermID get term info by term_id
func (t *Term) GetByID(db *gorm.DB) error {
	return db.First(&t, t.ID).Error
}

// GetTermTaxonomyByTermID get taxonomy terms info by term_id
func (t *TermTaxonomy) GetTermTaxonomyByTermID(db *gorm.DB) error {
	if err := db.Where("term_id = ?", t.TermID).First(t).Error; err != nil {
		return err
	}

	if err := db.Model(t).Association("Term").Find(&t.Term); err != nil {
		return err
	}

	return nil
}

func (t *Term) Save(db *gorm.DB) error {
	return db.Save(t).Error
}

func (t *Term) Update(db *gorm.DB, columns ...string) error {
	return db.Model(t).Select(columns).Updates(*t).Error
}

func (t *Term) Delete(db *gorm.DB) error {
	return db.Delete(t).Error
}

func (t *TermTaxonomy) Save(db *gorm.DB) error {
	return db.Save(t).Error
}

func (t *TermTaxonomy) Update(db *gorm.DB, columns ...string) error {
	return db.Model(t).Select(columns).Updates(*t).Error
}

func (t *TermTaxonomy) Delete(db *gorm.DB) error {
	return db.Delete(t).Error
}

// GetAllByType gets terms and taxonomy_terms by type(category, tag)
func (t *TermTaxonomy) GetAllByType(db *gorm.DB, taxomonyType string) ([]*TermTaxonomy, error) {
	var termTaxonomys []*TermTaxonomy
	err := db.Where("`taxonomy` = ?", taxomonyType).Preload("Term").Find(&termTaxonomys).Error
	return termTaxonomys, err
}

// TaxonomyCheckNameExist check the taxonomy name if is already exist
func CheckTaxonomyNameExist(db *gorm.DB, name, taxonomy string) bool {
	var count int64 = 0
	db.Table("pt_term AS t").
		Select("t.term_id, t.name").
		Joins("INNER JOIN pt_term_taxonomy AS tt ON tt.term_id = t.term_id").
		Where("t.name = ? AND tt.taxonomy = ?", name, taxonomy).
		Count(&count)

	if count > 0 {
		return true
	}

	return false
}
