package dao

import (
	"github.com/puti-projects/puti/internal/model"
	"gorm.io/gorm"
)

// GetTaxonomyLevel get term taxonomy level
// category returns actual level
// tag returns 0
func (d *Dao) GetTaxonomyLevel(termID uint64, taxonomyType string) (level uint64, err error) {
	// category
	if taxonomyType == "category" && termID != 0 {
		termTaxonomy := &model.TermTaxonomy{
			TermID:   termID,
			Taxonomy: taxonomyType,
		}
		err = termTaxonomy.GetByTermID(d.db)
		if err != nil {
			return
		}
		level = termTaxonomy.Level
		return
	}

	// tag
	return
}

func (d *Dao) CreateTaxonomy(tt *model.TermTaxonomy) error {
	return tt.Create(d.db)
}

func (d *Dao) GetTaxonomyByTermID(termID uint64, taxonomyType string) (*model.TermTaxonomy, error) {
	termTaxonomy := &model.TermTaxonomy{}
	termTaxonomy.TermID = termID
	if "all" != taxonomyType {
		termTaxonomy.Taxonomy = taxonomyType
	}

	if err := termTaxonomy.GetByTermID(d.db); err != nil {
		return nil, err
	}

	return termTaxonomy, nil
}

func (d *Dao) GetTermTaxonomyByTermID(termID uint64) (*model.TermTaxonomy, error) {
	termTaxonomy := &model.TermTaxonomy{
		TermID: termID,
	}
	if err := termTaxonomy.GetTermTaxonomyByTermID(d.db); err != nil {
		return nil, err
	}

	return termTaxonomy, nil
}

// UpdateTaxonomy update term and term taxonomy
// transcation is in use
// reset the parent's count number and child's level
func (d *Dao) UpdateTaxonomy(termTaxonomy *model.TermTaxonomy, taxonomyType string) error {
	// begin transcation
	tx := d.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// update term
	oldTerm := &model.Term{
		ID:          termTaxonomy.TermID,
		Name:        termTaxonomy.Term.Name,
		Slug:        termTaxonomy.Term.Slug,
		Description: termTaxonomy.Term.Description,
	}
	if err := oldTerm.Update(tx, "Name", "Slug", "Description"); err != nil {
		tx.Rollback()
		return err
	}

	// update term taxonomy (if change)
	// get old term taxonomy
	oldTermTaxonomy := &model.TermTaxonomy{
		TermID:   termTaxonomy.TermID,
		Taxonomy: taxonomyType,
	}
	if err := oldTermTaxonomy.GetByTermID(tx); err != nil {
		return err
	}

	// if changed parent term id
	// should update all related counting
	if oldTermTaxonomy.ParentTermID != termTaxonomy.ParentTermID {
		oldParentID := oldTermTaxonomy.ParentTermID // old parent termID

		// set new parent termID
		oldTermTaxonomy.ParentTermID = termTaxonomy.ParentTermID

		// get new parent's taxonomy level; tag will be 1 always
		// set new level
		var newLevel uint64
		if termTaxonomy.ParentTermID == 0 {
			// if new parentID = 0
			newLevel = 1
		} else {
			parentTermTaxonomy := &model.TermTaxonomy{TermID: termTaxonomy.ParentTermID}
			err := parentTermTaxonomy.GetColumnByTermID(tx, "level")
			if err != nil {
				tx.Rollback()
				return err
			}
			level := parentTermTaxonomy.Level
			newLevel = level + 1
		}

		ifChangeLevel := false
		if oldTermTaxonomy.Level != newLevel {
			ifChangeLevel = true
		}
		oldTermTaxonomy.Level = newLevel

		// save changed
		if err := oldTermTaxonomy.Update(tx, "ParentTermID", "Level"); err != nil {
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
			if err := updateTaxonomyChildLevel(tx, termTaxonomy.TermID, newLevel, taxonomyType); err != nil {
				tx.Rollback()
				return err
			}
		}

	}

	// commit
	return tx.Commit().Error
}

// updateTaxonomyChildLevel update category's all children level
func updateTaxonomyChildLevel(tx *gorm.DB, termID uint64, parentLevel uint64, taxonomyType string) error {
	// update all children
	t := &model.TermTaxonomy{}
	if err := tx.Table(t.TableName()).
		Where("parent_term_id = ?", termID).
		Updates(map[string]interface{}{"level": parentLevel + 1}).
		Error; err != nil {
		return err
	}

	// check children's child
	where := "`parent_term_id` = ? AND `taxonomy` = ?"
	whereArgs := []interface{}{termID, taxonomyType}
	children, err := t.Get(tx, where, whereArgs)
	if err != nil {
		return err
	}

	for _, item := range children {
		return updateTaxonomyChildLevel(tx, item.ID, parentLevel+1, taxonomyType)
	}

	return nil
}

// updateTaxonomyParentCount function of update count number
func updateTaxonomyParentCount(tx *gorm.DB, oldParentID, newParentID, countDiff uint64, taxonomyType string) error {
	// old parent's count
	// -countDiff
	if oldParentID != 0 {
		if err := updateTaxonomyCount(tx, oldParentID, -countDiff, taxonomyType); err != nil {
			return err
		}
	}

	// new parent's count
	// +countDiff
	if newParentID != 0 {
		if err := updateTaxonomyCount(tx, newParentID, countDiff, taxonomyType); err != nil {
			return err
		}
	}

	return nil
}

// GetTermChildrenNumber count children number
func (d *Dao) GetTermChildrenNumber(termID uint64, taxonomyType string) (count int64, err error) {
	if taxonomyType != "category" {
		return
	}
	t := &model.TermTaxonomy{}
	count, err = t.Count(d.db, "`parent_term_id` = ? AND `taxonomy` = ?", []interface{}{termID, taxonomyType})
	return
}

// DeleteTaxonomy delete term directly
func (d *Dao) DeleteTaxonomy(termID uint64, taxonomyType string) error {
	// begin transcation
	tx := d.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// get term count (taxonomy's article number)
	term := &model.Term{
		ID: termID,
	}
	if err := term.GetByID(tx); err != nil {
		tx.Rollback()
		return err
	}
	deletedCount := term.Count

	// delete term
	if err := term.Delete(tx); err != nil {
		tx.Rollback()
		return err
	}

	// get term_taxonomy id (user for delete relationship)
	termTaxonomy := &model.TermTaxonomy{
		TermID:   termID,
		Taxonomy: taxonomyType,
	}
	if err := termTaxonomy.GetByTermID(tx); err != nil {
		tx.Rollback()
		return err
	}
	termTaxonomyID := termTaxonomy.ID // Note: term_taxonomy_id
	parentID := termTaxonomy.ParentTermID

	// delete term taxonomy
	if err := termTaxonomy.Delete(tx); err != nil {
		tx.Rollback()
		return err
	}

	// delete term relationship with articles
	termRelationships := model.TermRelationships{}
	if err := termRelationships.DeleteByCondition(tx, "`term_taxonomy_id` = ?", []interface{}{termTaxonomyID}); err != nil {
		tx.Rollback()
		return err
	}

	// update parent's count number(if has parent)
	if parentID > 0 {
		if err := updateTaxonomyCount(tx, parentID, -deletedCount, taxonomyType); err != nil {
			tx.Rollback()
			return err
		}
	}

	// commit
	return tx.Commit().Error
}

// GetAllByType get all taxonomy by type
func (d *Dao) GetAllByType(taxonomyType string) ([]*model.TermTaxonomy, error) {
	t := &model.TermTaxonomy{}
	return t.GetAllByType(d.db, taxonomyType)
}

// CheckTaxonomyNameExist check if taxonomy name are aleady exist
func (d *Dao) CheckTaxonomyNameExist(name, taxonomy string) bool {
	return model.CheckTaxonomyNameExist(d.db, name, taxonomy)
}

// updateTaxonomyCount update count number using single connection
func updateTaxonomyCount(tx *gorm.DB, termID uint64, countDiff uint64, taxonomyType string) error {
	// get and update term
	Term := &model.Term{ID: termID}
	if err := Term.GetByID(tx); err != nil {
		return err
	}
	Term.Count = Term.Count + countDiff
	if err := Term.Save(tx); err != nil {
		return err
	}

	TermTaxonomy := &model.TermTaxonomy{
		TermID:   termID,
		Taxonomy: taxonomyType,
	}
	if err := TermTaxonomy.GetByTermID(tx); err != nil {
		return err
	}

	// if it has parent, update its parent's count
	if TermTaxonomy.ParentTermID != 0 {
		return updateTaxonomyCount(tx, TermTaxonomy.ParentTermID, countDiff, taxonomyType)
	}

	return nil
}

// updateTaxonomyCountByArticleChange update taxonomy article number count when editing the article
// diffenet with ”updateTaxonomyCount“ function (Direct calculation by level)
// Note: checkout taxonomy's parent and compare it with the term group is in need
// Note: termIDGroup 的所有 ID 中可能他们的父 ID 中存在在 termIDGroup 中（因为选择分类时可以随意勾选，不严格按照层级）
// 这里进行检测并去重
func updateTaxonomyCountByArticleChange(tx *gorm.DB, termIDGroup []uint64, countDiff int64) (err error) {
	var parentIDGroup []uint64
	for _, termID := range termIDGroup {
		parentIDGroup, err = getTaxonomyParentTermID(tx, termID, parentIDGroup)
		if err != nil {
			return err
		}
	}

	// If parentID not in termIDGroup, put it into termIDGroup
	if len(parentIDGroup) != 0 {
		for _, v := range parentIDGroup {
			inGroup := false
			for _, vv := range termIDGroup {
				if vv == v {
					inGroup = true
				}
			}

			if inGroup == false {
				termIDGroup = append(termIDGroup, v)
			}
		}
	}

	// 批量 +countDiff； 创建文章时传入的 countDiff 为 1
	if len(termIDGroup) != 0 {
		err = tx.Model(&model.Term{}).Where("term_id IN (?)", termIDGroup).UpdateColumn("count", gorm.Expr("count + ?", countDiff)).Error
		if err != nil {
			return err
		}
	}

	return nil
}

// getTaxonomyParentTermID find and get all parentID level by level
// return a slice include all level's parent ID
func getTaxonomyParentTermID(tx *gorm.DB, termID uint64, parentTerm []uint64) (parentTermGroup []uint64, err error) {
	termTaxonomy := &model.TermTaxonomy{
		TermID: termID,
	}
	err = termTaxonomy.GetByTermID(tx)
	if err != nil {
		return nil, err
	}

	// if it has parent
	if termTaxonomy.ParentTermID != 0 {
		parentTermGroup := append(parentTermGroup, termTaxonomy.ParentTermID)
		parentTermGroup, err := getTaxonomyParentTermID(tx, termTaxonomy.ParentTermID, parentTermGroup)
		if err != nil {
			return nil, err
		}
	}

	return parentTermGroup, nil
}
