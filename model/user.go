package model

import (
	"puti/pkg/auth"
	"puti/pkg/constvar"
)

// UserModel user model
type UserModel struct {
	Model

	Username string `gorm:"column:account;unique;not null"  validate:"min=1,max=60"`
	Password string `gorm:"column:password;not null" validate:"min=6,max=255"`
	Nickname string `gorm:"column:nickname;not null"`
	Email    string `gorm:"column:email;unique"`
	Avatar   string `gorm:"column:avatar"`
	PageURL  string `gorm:"column:page_url"`
	Status   int    `gorm:"column:status" sql:"index"`
	Roles    string `gorm:"column:role"`
}

// TableName is the user table name in db
func (c *UserModel) TableName() string {
	return "pt_users"
}

// Create creates a new user account
func (c *UserModel) Create() error {
	return DB.Local.Create(&c).Error
}

// Encrypt the user password.
func (c *UserModel) Encrypt() (err error) {
	c.Password, err = auth.Encrypt(c.Password)
	return
}

// GetUser gets an user by the user identifier.
func GetUser(username string) (*UserModel, error) {
	u := &UserModel{}
	d := DB.Local.Where("account = ?", username).First(&u)
	return u, d.Error
}

// DeleteUser deletes the user by id
func DeleteUser(id uint64) error {
	user := UserModel{}
	user.ID = id
	return DB.Local.Delete(&user).Error
}

// Update updates an user account information.
func (c *UserModel) Update() error {
	var user UserModel
	DB.Local.First(&user, c.ID)
	return DB.Local.Model(user).Save(c).Error
}

// Validate checks the update params
// TODO
func (c *UserModel) Validate() error {
	return nil
}

// ListUser List all users
func ListUser(username, role string, number, page, status int) ([]*UserModel, uint64, error) {
	if number == 0 {
		number = constvar.DefaultLimit
	}

	users := make([]*UserModel, 0)
	var count uint64

	where := "1"
	whereArgs := []interface{}{}
	if username != "" {
		where += " AND `nickname` LIKE ?"
		whereArgs = append(whereArgs, "%"+username+"%")
	}

	if role != "" {
		where += " AND `role` = ?"
		whereArgs = append(whereArgs, role)
	}

	if status == 0 || status == 1 {
		where += " AND `status` = ?"
		whereArgs = append(whereArgs, status)
	}

	if err := DB.Local.Model(&UserModel{}).Where(where, whereArgs...).Count(&count).Error; err != nil {
		return users, count, err
	}

	offset := (page - 1) * number
	if err := DB.Local.Where(where, whereArgs...).Offset(offset).Limit(number).Order("id desc").Find(&users).Error; err != nil {
		return users, count, err
	}

	return users, count, nil
}
