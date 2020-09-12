package model

import (
	"github.com/puti-projects/puti/internal/pkg/auth"

	"gorm.io/gorm"
)

// User user model
type User struct {
	Model

	Username string `gorm:"column:account;unique;not null"  validate:"min=1,max=60"`
	Password string `gorm:"column:password;not null" validate:"min=6,max=255"`
	Nickname string `gorm:"column:nickname;not null"`
	Email    string `gorm:"column:email;unique"`
	Avatar   string `gorm:"column:avatar"`
	PageURL  string `gorm:"column:page_url"`
	Status   int    `gorm:"column:status;default:1" sql:"index"`
	Roles    string `gorm:"column:role;default:subscriber"`
}

// TableName is the user table name in db
func (u *User) TableName() string {
	return "pt_user"
}

// Create creates a new user account
func (u *User) Create(db *gorm.DB) error {
	return db.Create(u).Error
}

// Encrypt the user password.
func (u *User) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return err
}

// Get get an user by the user identifier
func (u *User) Get(db *gorm.DB) error {
	if u.Username != "" {
		db = db.Where("status = 1 AND deleted_time is null AND account = ?", u.Username)
	}
	return db.First(u).Error
}

// GetByID get user by ID
func (u *User) GetByID(db *gorm.DB) error {
	return db.First(u, u.ID).Error
}

// Delete delete a user by id
func (u *User) Delete(db *gorm.DB) error {
	return db.Delete(u).Error
}

// Update updates an user account information
func (u *User) Update(db *gorm.DB) error {
	return db.Save(u).Error
}

// Count count user
func (u *User) Count(db *gorm.DB, where string, whereArgs []interface{}) (int64, error) {
	var count int64
	err := db.Model(u).Where(where, whereArgs...).Count(&count).Error
	return count, err
}

// List get user list
func (u *User) List(db *gorm.DB, where string, whereArgs []interface{}, offset, number int) ([]*User, error) {
	users := make([]*User, 0)
	err := db.Where(where, whereArgs...).Offset(offset).Limit(number).Order("id desc").Find(&users).Error
	return users, err
}
