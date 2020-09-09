package model

import (
	"time"

	"gorm.io/gorm"
)

// Model base model
type Model struct {
	ID        uint64         `gorm:"primaryKey;autoIncrement;column:id"`
	CreatedAt time.Time      `gorm:"column:created_time;not null"`
	UpdatedAt time.Time      `gorm:"column:updated_time;not null"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_time;DEFAULT null" sql:"index"`
}
