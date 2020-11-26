package dao

import (
	"time"

	"github.com/puti-projects/puti/internal/model"
	"github.com/puti-projects/puti/internal/pkg/db"
)

// KnowledgeResult
type KnowledgeResult struct {
	ID            uint64
	Name          string
	Slug          string
	Type          string
	Description   string
	CoverImageURL string
	UpdatedTime   time.Time
}

// GetKnowledgeList get knowledge list
func GetKnowledgeList() ([]*KnowledgeResult, error) {
	var result []*KnowledgeResult
	if err := db.Engine.Model(&model.Knowledge{}).
		Select("pt_knowledge.`id`, pt_knowledge.`name`, pt_knowledge.`slug`, pt_knowledge.`type`, pt_knowledge.`description`," +
			"pt_knowledge.`updated_time`, pt_resource.`guid` as cover_image_url").
		Joins("LEFT JOIN pt_resource ON pt_resource.id = pt_knowledge.`cover_image`").
		Where("pt_knowledge.`deleted_time` is null").
		Order("pt_knowledge.`updated_time` desc").
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

// KnowledgeItemResult
type KnowledgeItemResult struct {
	ID             uint64
	Symbol         uint64
	Title          string
	ContentVersion uint64
	ParentID       uint64
	Level          uint64
	Index          int64
}

// GetKnowledgeItemList get knowledge item list
func GetKnowledgeItemList(knowledgeID uint64) ([]*KnowledgeItemResult, error) {
	var result []*KnowledgeItemResult
	if err := db.Engine.Model(&model.KnowledgeItem{}).
		Select("`id`, `symbol`, `title`, `content_version`, `parent_id`, `level`, `index`").
		Where("`knowledge_id` = ?", knowledgeID).
		Order("`index` asc").
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

// GetKnowledgeBySlug get knowledge by knowledge slug and type
func GetKnowledgeBySlug(kType, kSlug string) (*model.Knowledge, error) {
	k := &model.Knowledge{}
	if err := db.Engine.Model(&model.Knowledge{}).Where("`slug` = ? AND `type` = ?", kSlug, kType).First(k).
		Error; err != nil {
		return nil, err
	}
	return k, nil
}

// KnowledgeItemContentResult
type KnowledgeItemContentResult struct {
	Symbol  uint64
	Title   string
	Content string
}

// GetKnowledgeItemContentBySymbol get knowledge item content by knowledge item symbol
func GetKnowledgeItemContentBySymbol(kiSymbol string) (*KnowledgeItemContentResult, error) {
	result := &KnowledgeItemContentResult{}
	if err := db.Engine.Model(&model.KnowledgeItem{}).
		Select("pt_knowledge_item.`symbol`, pt_knowledge_item.`title`, pt_knowledge_item_content.`content`").
		Joins("INNER JOIN pt_knowledge_item_content ON pt_knowledge_item_content.version = pt_knowledge_item.`content_version`").
		Where("pt_knowledge_item.`symbol` = ? AND pt_knowledge_item_content.`status` = ?", kiSymbol, 1).
		First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
