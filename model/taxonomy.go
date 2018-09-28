package model

// TermTaxonomyModel `pt_term_taxonomy`'s struct using GORM Has-One model
type TermTaxonomyModel struct {
	ID           uint64    `gorm:"column:term_taxonomy_id;not null;primary_key"`
	Term         TermModel `gorm:"foreignkey:TermID;association_foreignkey:ID"`
	TermID       uint64    `gorm:"column:term_id;not null"`
	ParentTermID uint64    `gorm:"column:parent_term_id;not null"`
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

// TableName TermTaxonomyModel's binding db name
func (c *TermTaxonomyModel) TableName() string {
	return "pt_term_taxonomy"
}

// TableName TermModel's binding db name
func (c *TermModel) TableName() string {
	return "pt_term"
}

// GetAllTermsByType gets terms and taxonomy_terms by type(category, tag)
func GetAllTermsByType(taxomonyType string) ([]*TermTaxonomyModel, error) {
	var termTaxonomys []*TermTaxonomyModel
	result := DB.Local.Where("taxonomy = ?", taxomonyType).Preload("Term").Find(&termTaxonomys)

	return termTaxonomys, result.Error
}

// GetTermsInfo get taxonomy terms info by term_id
func GetTermsInfo(termID uint64) (*TermTaxonomyModel, error) {
	termTaxonomy := &TermTaxonomyModel{}

	model := DB.Local.Where("term_id = ?", termID).First(&termTaxonomy)
	if model.Error != nil {
		return nil, model.Error
	}

	result := DB.Local.Model(&termTaxonomy).Related(&termTaxonomy.Term, "TermID")
	return termTaxonomy, result.Error
}
