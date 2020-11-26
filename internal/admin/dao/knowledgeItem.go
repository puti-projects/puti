package dao

import (
	"github.com/puti-projects/puti/internal/model"

	"gorm.io/gorm"
)

// CreateKnowledgeItem create knowledge item by association mode
func (d *Dao) CreateKnowledgeItem(kItem *model.KnowledgeItem) (*model.KnowledgeItem, error) {
	err := d.db.Transaction(func(tx *gorm.DB) error {
		// create
		if err := kItem.Create(tx); err != nil {
			return err
		}

		// update others index
		err := tx.Model(&model.KnowledgeItem{}).
			Where("`id` != ? AND `knowledge_id` = ? AND `level` = ?", kItem.ID, kItem.KnowledgeID, kItem.Level).
			Update("index", gorm.Expr("`index` + ?", 1)).Error
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return kItem, err
}

// GetKnowledgeItemListByKnowledgeID get all knowledge items by knowledge id
func (d *Dao) GetKnowledgeItemListByKnowledgeID(kID uint64) ([]*model.KnowledgeItem, error) {
	kItem := &model.KnowledgeItem{
		KnowledgeID: kID,
	}
	return kItem.GetAll(d.db)
}

// GetKnowledgeItemByID get knowledge item by ID
func (d *Dao) GetKnowledgeItemByID(kiID uint64) (*model.KnowledgeItem, error) {
	kItem := &model.KnowledgeItem{
		Model: model.Model{ID: kiID},
	}
	err := kItem.GetByID(d.db)
	if err != nil {
		return nil, err
	}

	return kItem, nil
}

// GetKnowledgeIDByItemID get knowledge item's knowledge ID
func (d *Dao) GetKnowledgeIDByItemID(kiID uint64) (uint64, error) {
	kItem, err := d.GetKnowledgeItemByID(kiID)
	if err != nil {
		return 0, err
	}
	return kItem.KnowledgeID, err
}

// GetKnowledgeItemSymbolByID get knowledge item symbol (unique sign) by ID
func (d *Dao) GetKnowledgeItemSymbolByID(kiID uint64) (uint64, error) {
	kItem, err := d.GetKnowledgeItemByID(kiID)
	if err != nil {
		return 0, err
	}
	return kItem.Symbol, err
}

// GetKnowledgeItemLastContent get last version of item content
func (d *Dao) GetKnowledgeItemLastContent(ki *model.KnowledgeItem) (*model.KnowledgeItemContent, error) {
	itemContents, err := ki.GetItemContent(d.db, "", "updated_time desc")
	if err != nil {
		return nil, err
	}
	return itemContents[0], err
}

// GetKnowledgeItemContentByVersion get knowledge item by version
func (d *Dao) GetKnowledgeItemContentByVersion(kItemID, version uint64) (*model.KnowledgeItemContent, error) {
	kic := &model.KnowledgeItemContent{
		KnowledgeItemID: kItemID,
		Version:         version,
	}
	res, err := kic.GetByVersion(d.db)
	return res, err
}

// CreateKnowledgeItemContent create a new version knowledge item content
func (d *Dao) CreateKnowledgeItemContent(kItemContent *model.KnowledgeItemContent) error {
	return kItemContent.Create(d.db)
}

// UpdateKnowledgeItemWithNodeChange update knowledge item with node change
func (d *Dao) UpdateKnowledgeItemWithNodeChange(kItemID uint64, newParentID uint64, indexChange string, indexRelatedNode uint64) error {
	return d.db.Transaction(func(tx *gorm.DB) error {
		kItem := &model.KnowledgeItem{
			Model: model.Model{ID: kItemID},
		}
		if err := kItem.GetByID(tx); err != nil {
			return err
		}

		// if change tree node
		if newParentID != kItem.ParentID {
			kItem.ParentID = newParentID

			// if set to root
			if newParentID == 0 {
				kItem.Level = 1
			} else {
				// get new parent
				kItemParent := &model.KnowledgeItem{
					Model: model.Model{ID: newParentID},
				}
				if err := kItemParent.GetByID(tx); err != nil {
					return err
				}
				// set level by parent's level
				kItem.Level = kItemParent.Level + 1
			}

			// update item
			if err := kItem.Save(tx); err != nil {
				return err
			}

			// update item's children
			if err := updateChildrenLevel(tx, kItemID, kItem.Level, kItem.KnowledgeID); err != nil {
				return err
			}
		}

		// get related item
		relatedItem := &model.KnowledgeItem{
			Model: model.Model{ID: indexRelatedNode},
		}
		if err := relatedItem.GetByID(tx); err != nil {
			return err
		}

		if indexChange == "before" {
			kItem.Index = relatedItem.Index - 1
		} else if indexChange == "after" {
			kItem.Index = relatedItem.Index + 1
		} else if indexChange == "inner" {
			kItem.Index = 0
		}
		// update index in this level
		if err := updateAllIndexInLevel(tx, relatedItem, indexChange); err != nil {
			return err
		}

		// update item (should after last step )
		if err := kItem.Save(tx); err != nil {
			return err
		}
		return nil
	})
}

func updateAllIndexInLevel(tx *gorm.DB, relatedItem *model.KnowledgeItem, indexChange string) error {
	sql := tx.Model(&model.KnowledgeItem{})
	if indexChange == "before" {
		sql.Where("`knowledge_id` = ? AND `parent_id` = ? AND `level` = ? AND `index` <= ?", relatedItem.KnowledgeID, relatedItem.ParentID, relatedItem.Level, relatedItem.Index-1).
			Update("index", gorm.Expr("`index` - ?", 1))
	} else if indexChange == "after" {
		sql.Where("`knowledge_id` = ? AND `parent_id` = ? AND `level` = ? AND `index` >= ?", relatedItem.KnowledgeID, relatedItem.ParentID, relatedItem.Level, relatedItem.Index+1).
			Update("index", gorm.Expr("`index` + ?", 1))
	}
	if err := sql.Error; err != nil {
		return err
	}
	return nil
}

func updateChildrenLevel(tx *gorm.DB, kItemID uint64, parentLevel uint64, knowledgeID uint64) error {
	k := &model.KnowledgeItem{}

	// check children
	where := "`knowledge_id` = ? AND `parent_id` = ?"
	whereArgs := []interface{}{knowledgeID, kItemID}
	children, rowsAffected, err := k.Get(tx, where, whereArgs)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return nil
	}

	// update all children
	if err := tx.Table(k.TableName()).
		Where("`knowledge_id` = ? AND parent_id = ?", knowledgeID, kItemID).
		Updates(map[string]interface{}{"level": parentLevel + 1}).
		Error; err != nil {
		return err
	}

	// check children's children
	for _, chill := range children {
		return updateChildrenLevel(tx, chill.ID, parentLevel+1, knowledgeID)
	}

	return nil
}

// UpdateKnowledgeItemTitle update knowledge item title
func (d *Dao) UpdateKnowledgeItemTitle(kItemID uint64, title string) error {
	kItem, err := d.GetKnowledgeItemByID(kItemID)
	if err != nil {
		return err
	}
	kItem.Title = title
	return kItem.Save(d.db)
}

// UpdateKnowledgeItem update knowledge item
func (d *Dao) UpdateKnowledgeItem(kItem *model.KnowledgeItem) error {
	return kItem.Save(d.db)
}

// UpdateKnowledgeItemContent update knowledge item content
func (d *Dao) UpdateKnowledgeItemContent(kItemContent *model.KnowledgeItemContent) error {
	return kItemContent.Save(d.db)
}

// ChangePublishedKnowledgeItemContent change current published version of content
func (d *Dao) ChangePublishedKnowledgeItemContent(kItemContent *model.KnowledgeItemContent) error {
	// get item info
	kItem := &model.KnowledgeItem{
		Model: model.Model{ID: kItemContent.KnowledgeItemID},
	}
	if err := kItem.GetByID(d.db); err != nil {
		return err
	}

	var nowKItemContent []model.KnowledgeItemContent
	if err := d.db.Model(&kItem).Association("ItemContents").Find(&nowKItemContent); err != nil {
		return err
	}

	// bug
	for k, v := range nowKItemContent {
		if v.Version == kItemContent.Version {
			nowKItemContent[k].Status = model.KnowledgeItemContentStatusCurrent
			nowKItemContent[k].Content = kItemContent.Content
		}

		// if published before; old version set to 0
		if kItem.ContentVersion > 0 {
			if v.Status == model.KnowledgeItemContentStatusCurrent {
				nowKItemContent[k].Status = model.KnowledgeItemContentStatusCommon
			}
		}
	}

	// update current version in item info
	kItem.ContentVersion = kItemContent.Version
	kItem.ItemContents = nowKItemContent
	if err := d.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&kItem).Error; err != nil {
		return err
	}

	return nil
}

// CheckKnowledgeItemHasChildren check whether the knowledge item has children
func (d *Dao) CheckKnowledgeItemHasChildren(kItemID, knowledgeID uint64) (bool, error) {
	var count int64
	err := d.db.Model(&model.KnowledgeItem{}).Where("`knowledge_id` = ? AND `parent_id` = ?", knowledgeID, kItemID).Count(&count).Error
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}
	return false, nil
}

// DeleteKnowledgeItem delete knowledge item
func (d *Dao) DeleteKnowledgeItem(kItemID uint64) error {
	kItem := model.KnowledgeItem{
		Model: model.Model{ID: kItemID},
	}
	if err := kItem.Delete(d.db); err != nil {
		return err
	}
	return nil
}
