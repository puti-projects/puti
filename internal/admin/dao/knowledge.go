package dao

import (
	"database/sql"
	"time"

	"github.com/puti-projects/puti/internal/model"
)

// KnowledgeInfo knowledge base info
type KnowledgeInfo struct {
	ID             uint64
	Name           string
	Slug           string
	Type           string
	Description    string
	CoverImageID   uint64
	CoverImageName string
	CoverImageURL  string
	LastUpdated    sql.NullTime
	CreatedTime    time.Time
}

// CreateSubject create a knowledge base
func (d *Dao) CreateKnowledge(k *model.Knowledge) error {
	return k.Create(d.db)
}

// UpdateKnowledge update knowledge base info
func (d *Dao) UpdateKnowledge(k *model.Knowledge) error {
	oldK := &model.Knowledge{
		Model: model.Model{ID: k.ID},
	}
	if err := oldK.GetByID(d.db); err != nil {
		return err
	}

	oldK.Name = k.Name
	oldK.Slug = k.Slug
	oldK.Description = k.Description
	oldK.CoverImage = k.CoverImage

	if err := oldK.Save(d.db); err != nil {
		return err
	}

	return nil
}

// SaveKnowledge save knowledge base info
func (d *Dao) SaveKnowledge(k *model.Knowledge) error {
	return k.Save(d.db)
}

// GetAllKnowledgeList get all knowledge bases list contains related data
func (d *Dao) GetAllKnowledgeList() ([]*KnowledgeInfo, error) {
	var results []*KnowledgeInfo
	err := d.db.Model(&model.Knowledge{}).
		Select("pt_knowledge.id, pt_knowledge.name, pt_knowledge.slug, pt_knowledge.type, pt_knowledge.description, pt_resource.id as cover_image_id, pt_resource.title as cover_image_name, pt_resource.guid as cover_image_url, pt_knowledge.last_updated, pt_knowledge.created_time").
		Joins("LEFT JOIN pt_resource ON pt_resource.`id` = pt_knowledge.`cover_image` AND pt_resource.`status` = 1 AND pt_resource.`deleted_time` is null").
		Where("pt_knowledge.`deleted_time` is null").
		Find(&results).Error
	return results, err
}

// GetKnowledgeByID get knowledge by ID
func (d *Dao) GetKnowledgeByID(kID uint64) (*model.Knowledge, error) {
	knowledge := &model.Knowledge{
		Model: model.Model{ID: kID},
	}
	err := knowledge.GetByID(d.db)
	if err != nil {
		return nil, err
	}

	return knowledge, nil
}
