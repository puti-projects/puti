package model

import (
	"fmt"
	"time"

	"gingob/pkg/auth"
	"gingob/pkg/constvar"
	"gingob/pkg/errno"
)

// User model
type UserModel struct {
	Id             uint64    `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	Nickname       string    `gorm:"column:nickname;not null" json:"-"`
	Email          string    `gorm:"column:email" json:"-"`
	RegisteredTime time.Time `gorm:"column:registered_time" json:"-"`
	Status         int       `gorm:"column:status" sql:"index" json:"-"`
	Role           string    `gorm:"column:role" json:"-"`

	Username string `json:"username" gorm:"column:account;not null" binding:"required" validate:"min=1,max=32"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`
}

// User table in db
func (c *UserModel) TableName() string {
	return "gb_users"
}

// create a new user account
func (u *UserModel) Create() error {
	return DB.Local.Create(&u).Error
}

// Encrypt the user password.
func (u *UserModel) Encrypt() (err error) {
	u.Password, err = auth.Encrypt(u.Password)
	return
}

// GetUser gets an user by the user identifier.
func GetUser(username string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Local.Where("account = ?", username).First(&u)
	return u, d.Error
}

func DeleteUser(id uint64) error {
	user := UserModel{}
	user.Id = id
	return DB.Local.Delete(&user).Error
}

// Update updates an user account information.
func (u *UserModel) Update() error {
	return DB.Local.Save(u).Error
}

// todo
func (u *UserModel) Validate() error {
	if u.Username == "" {
		return errno.New(errno.ErrValidation, nil).Add("username is empty.")
	}

	if u.Password == "" {
		return errno.New(errno.ErrValidation, nil).Add("password is empty.")
	}

	return nil
}

// ListUser List all users
func ListUser(username string, offset, limit int) ([]*UserModel, uint64, error) {
	if limit == 0 {
		limit = constvar.DefaultLimit
	}

	users := make([]*UserModel, 0)
	var count uint64

	where := fmt.Sprintf("account like '%%%s%%'", username)
	if err := DB.Local.Model(&UserModel{}).Where(where).Count(&count).Error; err != nil {
		return users, count, err
	}

	if err := DB.Local.Where(where).Offset(offset).Limit(limit).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}

	return users, count, nil
}
