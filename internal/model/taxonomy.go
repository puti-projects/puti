package model

import "github.com/puti-projects/puti/internal/pkg/db"

// TermTaxonomyModel `pt_term_taxonomy`'s struct using GORM Has-One model
type TermTaxonomyModel struct {
	ID           uint64    `gorm:"column:term_taxonomy_id;not null;primaryKey"`
	Term         TermModel `gorm:"foreignKey:TermID;references:ID"`
	TermID       uint64    `gorm:"column:term_id;not null"`
	ParentTermID uint64    `gorm:"column:parent_term_id;not null"`
	Level        uint64    `gorm:"column:level;not null"`
	Taxonomy     string    `gorm:"column:taxonomy;not null"`
	TermGroup    uint64    `gorm:"column:term_group;default:0;not null"`
}

// TermModel `pt_terms`'s struct for taxomony info
type TermModel struct {
	ID          uint64 `gorm:"column:term_id;not null;primary_key"`
	Name        string `gorm:"column:name;not null"`
	Slug        string `gorm:"column:slug;not null"`
	Description string `gorm:"column:description;not null"`
	Count       uint64 `gorm:"column:count;not null"`
}

// DefaultUnCategorizedID the default category's ID
const DefaultUnCategorizedID = 1

// TableName TermTaxonomyModel's binding db name
func (c *TermTaxonomyModel) TableName() string {
	return "pt_term_taxonomy"
}

// TableName TermModel's binding db name
func (c *TermModel) TableName() string {
	return "pt_term"
}

// Create creates a new taxonomy term
func (c *TermModel) Create() error {
	return db.DBEngine.Create(&c).Error
}

// Update updates the taxonomy term
func (c *TermModel) Update() error {
	return db.DBEngine.Model(&TermModel{}).Save(c).Error
}

// Create creates a new taxonomy
func (c *TermTaxonomyModel) Create() error {
	return db.DBEngine.Create(&c).Error
}

// Update updates the taxonomy term taxonomy
func (c *TermTaxonomyModel) Update() (err error) {
	return db.DBEngine.Model(&TermTaxonomyModel{}).Save(c).Error
}

// GetAllTermsByType gets terms and taxonomy_terms by type(category, tag)
func GetAllTermsByType(taxomonyType string) ([]*TermTaxonomyModel, error) {
	var termTaxonomys []*TermTaxonomyModel
	result := db.DBEngine.Where("taxonomy = ?", taxomonyType).Preload("Term").Find(&termTaxonomys)

	return termTaxonomys, result.Error
}

// GetTermsInfo get taxonomy terms info by term_id
func GetTermsInfo(termID uint64) (*TermTaxonomyModel, error) {
	termTaxonomy := &TermTaxonomyModel{}

	model := db.DBEngine.Where("term_id = ?", termID).First(&termTaxonomy)
	if model.Error != nil {
		return nil, model.Error
	}

	err := db.DBEngine.Model(&termTaxonomy).Association("Term").Find(&termTaxonomy.Term)
	return termTaxonomy, err
}

// TaxonomyCheckNameExist check the taxonomy name if is already exist
func TaxonomyCheckNameExist(name, taxonomy string) bool {
	var count int64 = 0
	db.DBEngine.Table("pt_term t").
		Select("t.term_id, t.name").
		Joins("inner join pt_term_taxonomy tt on tt.term_id = t.term_id").
		Where("t.name = ? AND tt.taxonomy = ?", name, taxonomy).
		Count(&count)

	if count > 0 {
		return true
	}

	return false
}

// GetTaxonomyLevel calculate the level
func GetTaxonomyLevel(parentID uint64, taxonomy string) (level uint64, err error) {
	if taxonomy == "category" && parentID != 0 {
		// get parent level
		termTaxonomy := &TermTaxonomyModel{}
		d := db.DBEngine.Where("term_id = ? AND taxonomy = 'category'", parentID).First(&termTaxonomy)

		return termTaxonomy.Level + 1, d.Error
	}

	return 1, nil
}

// GetTermByID get term info by term_id
func GetTermByID(termID uint64) (*TermModel, error) {
	m := &TermModel{}
	d := db.DBEngine.Where("term_id = ?", termID).First(&m)
	return m, d.Error
}

// GetTermTaxonomy get term taxonomy by term_id and taxonomy type
func GetTermTaxonomy(termID uint64, taxonomyType string) (*TermTaxonomyModel, error) {
	m := &TermTaxonomyModel{}
	if taxonomyType == "" {
		d := db.DBEngine.Where("term_id = ?", termID).First(&m)
		return m, d.Error
	}

	d := db.DBEngine.Where("term_id = ? AND taxonomy = ?", termID, taxonomyType).First(&m)
	return m, d.Error
}

// GetTermChildrenNumber calcuelate the total number of term's children
func GetTermChildrenNumber(termID uint64, taxonomyType string) (count int64) {
	if taxonomyType != "category" {
		count = 0
		return count
	}

	db.DBEngine.Model(&TermTaxonomyModel{}).Where("parent_term_id = ? AND taxonomy = ?", termID, taxonomyType).Count(&count)
	return count
}
