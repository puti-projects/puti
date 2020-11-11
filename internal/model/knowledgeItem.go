package model

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type KnowledgeItem struct {
	Model

	KnowledgeID    uint64       `gorm:"column:knowledge_id;not null"`
	UserID         uint64       `gorm:"column:user_id;not null"`
	Title          string       `gorm:"column:title;not null"`
	ContentVersion uint64       `gorm:"column:content_version;not null"`
	GUID           string       `gorm:"column:guid;not null"`
	ParentID       uint64       `gorm:"column:parent_id;not null;default:0"`
	Level          uint64       `gorm:"column:level;not null;default:0"`
	Index          uint64       `gorm:"column:index;not null;default:0"`
	CommentCount   uint64       `gorm:"column:comment_count;not null;default:0"`
	ViewCount      uint64       `gorm:"column:view_count;not null;default:0"`
	LastPublished  sql.NullTime `gorm:"column:last_published;default:null"`

	ItemContents []KnowledgeItemContent `gorm:"foreignKey:KnowledgeItemID"`
}

type KnowledgeItemContent struct {
	ID              uint64    `gorm:"primaryKey;autoIncrement;column:id"`
	KnowledgeItemID uint64    `gorm:"column:knowledge_item_id;not null"`
	Version         int64     `gorm:"column:version;not null"`
	Status          uint8     `gorm:"column:status;not null;default:0"`
	Content         string    `gorm:"column:content;not null"`
	UpdatedAt       time.Time `gorm:"column:updated_time;not null"`
}

// TableName is the knowledge item table name in db
func (ki *KnowledgeItem) TableName() string {
	return "pt_knowledge_item"
}

// TableName is the knowledge item table name in db
func (kic *KnowledgeItemContent) TableName() string {
	return "pt_knowledge_item_content"
}

// GetAll get all knowledge items by knowledge id
func (ki *KnowledgeItem) GetAll(db *gorm.DB) ([]*KnowledgeItem, error) {
	var knowledgeItems []*KnowledgeItem
	err := db.Where("`knowledge_id` = ?", ki.KnowledgeID).Find(&knowledgeItems).Error
	return knowledgeItems, err
}
