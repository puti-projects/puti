package service

import (
	"strconv"

	"puti/model"

	"github.com/jinzhu/gorm"
)

// TermInfo terms info
type TermInfo struct {
	ID          uint64 `json:"term_id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Pid         uint64 `json:"parent_term_id"`
	Level       uint64 `json:"level"`
}

// TreeNode TaxonomyTree's node struct
type TreeNode struct {
	ID           uint64      `json:"id"`
	Name         string      `json:"name"`
	Slug         string      `json:"slug"`
	Description  string      `json:"description"`
	Count        uint64      `json:"count"`
	TermID       uint64      `json:"term_id"`
	ParentTermID uint64      `json:"pid"`
	Level        uint64      `json:"level"`
	Children     []*TreeNode `json:"children"`
}

// GetTaxonomyList get taxonomy tree by type and return a tree structure
func GetTaxonomyList(taxonomyType string) (taxonomyTree []*TreeNode, err error) {
	termTaxonomy, err := model.GetAllTermsByType(taxonomyType)
	if err != nil {
		return nil, err
	}

	list := GetTaxonomyTree(termTaxonomy, 0)

	return list, nil
}

// GetTaxonomyTree return a taxonomy tree
func GetTaxonomyTree(termTaxonomy []*model.TermTaxonomyModel, pid uint64) []*TreeNode {
	var tree []*TreeNode

	for _, v := range termTaxonomy {
		if pid == v.ParentTermID {
			treeNode := TreeNode{
				ID:           v.ID,
				Name:         v.Term.Name,
				Slug:         v.Term.Slug,
				Description:  v.Term.Description,
				Count:        v.Term.Count,
				TermID:       v.TermID,
				ParentTermID: v.ParentTermID,
				Level:        v.Level,
			}
			treeNode.Children = GetTaxonomyTree(termTaxonomy, v.ID)
			tree = append(tree, &treeNode)
		}
	}

	return tree
}

// GetTaxonomyInfo get term info by term_id
func GetTaxonomyInfo(termID string) (*TermInfo, error) {
	ID, _ := strconv.Atoi(termID)

	info, err := model.GetTermsInfo(uint64(ID))
	if err != nil {
		return nil, err
	}

	termInfo := &TermInfo{
		ID:          info.ID,
		Name:        info.Term.Name,
		Slug:        info.Term.Slug,
		Description: info.Term.Description,
		Pid:         info.ParentTermID,
		Level:       info.Level,
	}

	return termInfo, nil
}

// UpdateTaxonomy update term and term taxonomy
// transcation is in use
// reset the parent's count number and child's level
func UpdateTaxonomy(term *model.TermModel, termTaxonomy *model.TermTaxonomyModel, taxonomyType string) (err error) {
	// begin transcation
	tx := model.DB.Local.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return err
	}

	// udapte term
	oldTerm, err := model.GetTermByID(term.ID)
	if err != nil {
		return err
	}
	oldTerm.Name = term.Name
	oldTerm.Slug = term.Slug
	oldTerm.Description = term.Description
	if err = tx.Model(&model.TermModel{}).Save(oldTerm).Error; err != nil {
		tx.Rollback()
		return err
	}

	// update term taxonomy (if change)
	oldTermTaxonomy, err := model.GetTermTaxonomy(term.ID, taxonomyType)
	if err != nil {
		return err
	}
	if oldTermTaxonomy.ParentTermID != termTaxonomy.ParentTermID {
		oldParentID := oldTermTaxonomy.ParentTermID
		oldTermTaxonomy.ParentTermID = termTaxonomy.ParentTermID
		// get new taxonomy level; tag will be 1 always
		newLevel, err := model.GetTaxonomyLevel(termTaxonomy.ParentTermID, "category")
		if err != nil {
			return err
		}
		ifChangeLevel := false
		if oldTermTaxonomy.Level != newLevel {
			oldTermTaxonomy.Level = newLevel
			ifChangeLevel = true
		}

		if err = tx.Model(&model.TermTaxonomyModel{}).Save(oldTermTaxonomy).Error; err != nil {
			tx.Rollback()
			return err
		}

		// update parent's count
		if err := updateTaxonomyParentCount(tx, oldParentID, termTaxonomy.ParentTermID, oldTerm.Count, taxonomyType); err != nil {
			tx.Rollback()
			return err
		}

		// udapte child's level
		if ifChangeLevel {
			if err := updateTaxonomyChildLevel(tx, term.ID, newLevel, taxonomyType); err != nil {
				tx.Rollback()
				return err
			}
		}

	}

	// commit
	return tx.Commit().Error
}

// updateTaxonomyChildLevel update category's children level
// TODO BUG here 最后层更新失败
func updateTaxonomyChildLevel(tx *gorm.DB, termID, newLevel uint64, taxonomyType string) error {
	TermTaxonomy := []model.TermTaxonomyModel{}
	tt := tx.Where("parent_term_id = ? AND taxonomy = ?", termID, taxonomyType).Find(&TermTaxonomy)
	if tt.Error != nil {
		return tt.Error
	}

	// update all children
	t := &model.TermTaxonomyModel{}
	tx.Table(t.TableName()).Where("parent_term_id = ?", termID).
		Updates(map[string]interface{}{"level": newLevel + 1})

	for _, item := range TermTaxonomy {
		return updateTaxonomyChildLevel(tx, item.ID, item.Level, taxonomyType)
	}

	return nil
}

// updateTaxonomyParentCount main function of update count number
func updateTaxonomyParentCount(tx *gorm.DB, oldParentID, newParentID, countDiff uint64, taxonomyType string) error {
	// old parents count
	if oldParentID != 0 {
		if err := updateTaxonomyCount(tx, oldParentID, -countDiff, taxonomyType); err != nil {
			return err
		}
	}

	// new parents count
	if newParentID != 0 {
		if err := updateTaxonomyCount(tx, newParentID, countDiff, taxonomyType); err != nil {
			return err
		}
	}

	return nil
}

// updateTaxonomyCount update count number using single connection
func updateTaxonomyCount(tx *gorm.DB, termID uint64, countDiff uint64, taxonomyType string) error {
	Term := &model.TermModel{}
	t := tx.Where("term_id = ?", termID).First(&Term)
	if t.Error != nil {
		return t.Error
	}
	Term.Count = Term.Count + countDiff
	if err := tx.Model(&model.TermModel{}).Save(Term).Error; err != nil {
		return err
	}

	TermTaxonomy := &model.TermTaxonomyModel{}
	tt := tx.Where("term_id = ? AND taxonomy = ?", termID, taxonomyType).First(&TermTaxonomy)
	if tt.Error != nil {
		return tt.Error
	}
	if TermTaxonomy.ParentTermID != 0 {
		return updateTaxonomyCount(tx, TermTaxonomy.ParentTermID, countDiff, taxonomyType)
	}

	return nil
}
