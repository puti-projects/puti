package model

import (
	"sync"
	"time"
)

// BaseModel base model
type BaseModel struct {
	ID             uint64     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	RegisteredTime time.Time  `gorm:"column:registered_time" json:"-"`
	Status         *time.Time `gorm:"column:status" sql:"index" json:"-"`
}

// UserInfo is the user struct for user list
type UserInfo struct {
	ID             uint64 `json:"id"`
	Accout         string `json:"account"`
	Nickname       string `json:"nickname"`
	Email          string `json:"email"`
	RegisteredTime string `json:"registered_time"`
	Roles          string `json:"roles"`
	Status         int    `json:"status"`
}

// UserList user list
type UserList struct {
	Lock  *sync.Mutex
	IDMap map[uint64]*UserInfo
}

// Token represents a JSON web token.
type Token struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
