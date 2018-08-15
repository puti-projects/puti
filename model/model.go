package model

import (
	"time"
)

// BaseModel base model
type BaseModel struct {
	ID             uint64     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	RegisteredTime time.Time  `gorm:"column:registered_time" json:"-"`
	Status         *time.Time `gorm:"column:status" sql:"index" json:"-"`
}

// Token represents a JSON web token.
type Token struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
