package model

import (
	"time"
)

// Model base model
type Model struct {
	ID        uint64     `gorm:"primary_key;AUTO_INCREMENT;column:id"`
	CreatedAt time.Time  `gorm:"column:created_time;not null"`
	UpdatedAt time.Time  `gorm:"column:updated_time;not null"`
	DeletedAt *time.Time `gorm:"column:deleted_time;DEFAULT null" sql:"index"`
}
