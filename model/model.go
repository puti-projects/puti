package model

import (
	"sync"
	"time"
)

type BaseModel struct {
	Id             uint64     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	RegisteredTime time.Time  `gorm:"column:registered_time" json:"-"`
	Status         *time.Time `gorm:"column:status" sql:"index" json:"-"`
}

type UserInfo struct {
	Id             uint64 `json:"id"`
	Accout         string `json:"account"`
	SayHello       string `json:"sayHello"`
	Nickname       string `json:"nickname"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	RegisteredTime string `json:"registered_time"`
	Role           string `json:"role"`
	Status         int    `json:"status"`
}

type UserList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*UserInfo
}

// Token represents a JSON web token.
type Token struct {
	Token string `json:"token"`
}
