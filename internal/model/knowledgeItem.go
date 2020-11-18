package model

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

// KnowledgeItemContentStatusCurrent knowledge item content current(published) version
const KnowledgeItemContentStatusCurrent = 1

// KnowledgeItemContentStatusCommon knowledge item content not current version
const KnowledgeItemContentStatusCommon = 0

type KnowledgeItem struct {
	Model

	KnowledgeID    uint64       `gorm:"column:knowledge_id;not null"`
	UserID         uint64       `gorm:"column:user_id;not null"`
	Title          string       `gorm:"column:title;not null"`
	ContentVersion uint64       `gorm:"column:content_version;not null"`
	GUID           string       `gorm:"column:guid;not null"`
	ParentID       uint64       `gorm:"column:parent_id;not null;default:0"`
	Level          uint64       `gorm:"column:level;not null;default:0"`
	Index          int64        `gorm:"column:index;not null;default:0"`
	CommentCount   uint64       `gorm:"column:comment_count;not null;default:0"`
	ViewCount      uint64       `gorm:"column:view_count;not null;default:0"`
	LastPublished  sql.NullTime `gorm:"column:last_published;default:null"`

	ItemContents []KnowledgeItemContent `gorm:"foreignKey:KnowledgeItemID"`
}

type KnowledgeItemContent struct {
	ID              uint64    `gorm:"primaryKey;autoIncrement;column:id"`
	KnowledgeItemID uint64    `gorm:"column:knowledge_item_id;not null"`
	Version         uint64    `gorm:"column:version;not null"`
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

// Create create a new knowledge item
func (ki *KnowledgeItem) Create(db *gorm.DB) error {
	return db.Create(ki).Error
}

// Create save the new version knowledge item content
func (kic *KnowledgeItemContent) Create(db *gorm.DB) error {
	return db.Create(kic).Error
}

// Save update knowledge item
func (ki *KnowledgeItem) Save(db *gorm.DB) error {
	return db.Save(ki).Error
}

// Save update knowledge item content
func (kic *KnowledgeItemContent) Save(db *gorm.DB) error {
	return db.Save(kic).Error
}

// GetByID get knowledge item by ID
func (ki *KnowledgeItem) GetByID(db *gorm.DB) error {
	return db.First(ki, ki.ID).Error
}

// Get get all by condition
func (ki *KnowledgeItem) Get(db *gorm.DB, where string, whereArgs []interface{}) ([]*KnowledgeItem, int64, error) {
	kItem := make([]*KnowledgeItem, 0)
	result := db.Where(where, whereArgs...).Find(&kItem)
	return kItem, result.RowsAffected, result.Error
}

// GetAll get all knowledge items by knowledge id
func (ki *KnowledgeItem) GetAll(db *gorm.DB) ([]*KnowledgeItem, error) {
	var knowledgeItems []*KnowledgeItem
	err := db.Where("`knowledge_id` = ?", ki.KnowledgeID).Order("`index` asc").Find(&knowledgeItems).Error
	return knowledgeItems, err
}

// GetByVersion get knowledge item by version
func (kic *KnowledgeItemContent) GetByVersion(db *gorm.DB) (*KnowledgeItemContent, error) {
	//var knowledgeItemContent *KnowledgeItemContent
	err := db.Where("knowledge_item_id = ? AND version = ?", kic.KnowledgeItemID, kic.Version).First(&kic).Error
	return kic, err
}

// GetLargeIndexInLevel get the largest index value in the same level
func (ki *KnowledgeItem) GetLargeIndexInLevel(db *gorm.DB) (uint64, error) {
	var kItem *KnowledgeItem
	if err := db.Select("index").Where("level = ?", ki.Level).Order("index desc").First(&kItem).Error; err != nil {
		return 0, err
	}
	return kItem.Level, nil
}

// GetLastContent get item contents
func (ki *KnowledgeItem) GetItemContent(db *gorm.DB, where string, order string) ([]*KnowledgeItemContent, error) {
	var itemContents []*KnowledgeItemContent
	query := db.Model(&ki)
	if where != "" {
		query = query.Where(where)
	}
	if order != "" {
		query = query.Order(order)
	}
	err := query.Association("ItemContents").Find(&itemContents)
	return itemContents, err
}

// Delete delete the knowledge item by id
func (ki *KnowledgeItem) Delete(db *gorm.DB) error {
	return db.Delete(ki).Error
}
