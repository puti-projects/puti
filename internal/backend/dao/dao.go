package dao

import (
	"github.com/puti-projects/puti/internal/pkg/db"

	"gorm.io/gorm"
)

type Dao struct {
	db *gorm.DB
}

var Engine *Dao

// NewDaoEngine create a dao instance
// Note: after db.DBEngine inited
func NewDaoEngine() {
	Engine = new(Dao)
	Engine.db = db.DBEngine
}
