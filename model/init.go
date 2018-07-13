package model

import "github.com/jinzhu/gorm"

type Database struct {
	Self   *gorm.DB
	Docker *gorm.DB
}
