package model

import "github.com/puti-projects/puti/internal/pkg/db"

// TermRelationshipsModel `pt_term_relationships` 's struct taxomony raltionships with object
type TermRelationshipsModel struct {
	ObjectID       uint64 `gorm:"column:object_id;not null;primary_key"`
	TermTaxonomyID uint64 `gorm:"column:term_taxonomy_id;not null;primary_key"`
	TermOrder      string `gorm:"column:term_order;not null"`
}

// ArticleTaxonomy use for article taxonomy
type ArticleTaxonomy struct {
	TermID   uint64
	Taxonomy string
}

// TableName TermRelationshipsModel's binding db name
func (c *TermRelationshipsModel) TableName() string {
	return "pt_term_relationships"
}

// GetArticleTaxonomy get article taxonomy include all type
func GetArticleTaxonomy(articleID uint64) ([]*ArticleTaxonomy, error) {
	sql := "SELECT t.term_id, tt.taxonomy FROM pt_term t LEFT JOIN pt_term_taxonomy tt ON tt.term_id = t.term_id LEFT JOIN pt_term_relationships tr ON tr.term_taxonomy_id = tt.term_taxonomy_id WHERE tr.object_id = ?"
	rows, err := db.DBEngine.Raw(sql, articleID).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]*ArticleTaxonomy, 0)
	for rows.Next() {
		articleTaxonomy := &ArticleTaxonomy{}

		if err := db.DBEngine.ScanRows(rows, &articleTaxonomy); err != nil {
			return nil, err
		}

		result = append(result, articleTaxonomy)
	}

	return result, nil
}
