package dao

import (
	"gorm.io/gorm"
)

// Dao data manipulation layer for frontend
type Dao struct {
	db *gorm.DB
}

// New return a pointer of a Dao instance
func New(db *gorm.DB) *Dao {
	return &Dao{
		db: db,
	}
}
