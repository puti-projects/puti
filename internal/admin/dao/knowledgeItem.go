package dao

import "github.com/puti-projects/puti/internal/model"

// CreateKnowledgeItem create knowledge item by association mode
func (d *Dao) CreateKnowledgeItem(kItem *model.KnowledgeItem) (*model.KnowledgeItem, error) {
	err := d.db.Create(kItem).Error
	return kItem, err
}

// GetKnowledgeItemListByKnowledgeID get all knowledge items by knowledge id
func (d *Dao) GetKnowledgeItemListByKnowledgeID(kID uint64) ([]*model.KnowledgeItem, error) {
	kItem := &model.KnowledgeItem{
		KnowledgeID: kID,
	}
	return kItem.GetAll(d.db)
}
