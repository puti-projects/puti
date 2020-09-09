package dao

import (
	"strings"

	"github.com/puti-projects/puti/internal/model"
	"github.com/puti-projects/puti/internal/pkg/constvar"
	"github.com/puti-projects/puti/internal/pkg/errno"
)

// GetUser get user by username
func (d *Dao) GetUser(username string) (*model.User, error) {
	user := &model.User{
		Username: username,
	}
	err := user.Get(d.db)
	return user, err
}

// GetUser get user by ID
func (d *Dao) GetUserByID(userID uint64) (*model.User, error) {
	user := &model.User{
		Model: model.Model{ID: userID},
	}
	err := user.GetByID(d.db)
	return user, err
}

// CreateUser create a new user
func (d *Dao) CreateUser(user *model.User) error {
	if err := user.Encrypt(); err != nil {
		return errno.New(errno.ErrEncrypt, err)
	}

	// Insert the user to the database.
	if err := user.Create(d.db); err != nil {
		return errno.New(errno.ErrDatabase, err)
	}

	return nil
}

// UpdateUser update user infomation
func (d *Dao) UpdateUser(userID uint64, nickname, email, website, role string) error {
	oldUser, err := d.GetUserByID(userID)
	if err != nil {
		return err
	}

	// Set new values
	nickname = strings.TrimSpace(nickname)
	if nickname == "" {
		// Set nickname to account if nickname is empty
		// username can not be change
		nickname = oldUser.Username
	}
	oldUser.Nickname = nickname
	oldUser.Email = strings.TrimSpace(email)
	oldUser.PageURL = strings.TrimSpace(website)
	oldUser.Roles = role

	err = oldUser.Update(d.db)
	if err != nil {
		return err
	}

	return nil
}

// UpdateUserStatus update user status
func (d *Dao) UpdateUserStatus(userID uint64, status int) error {
	oldUser, err := d.GetUserByID(userID)
	if err != nil {
		return err
	}

	// Set new status values
	oldUser.Status = status

	err = oldUser.Update(d.db)
	if err != nil {
		return err
	}

	return nil
}

// UpdateUserPassword reset user password by user ID
func (d *Dao) UpdateUserPassword(userID uint64, password string) error {
	oldUser, err := d.GetUserByID(userID)
	if err != nil {
		return err
	}

	// Set new password
	oldUser.Password = password

	// encrypt password
	if err := oldUser.Encrypt(); err != nil {
		return errno.New(errno.ErrEncrypt, err)
	}

	err = oldUser.Update(d.db)
	if err != nil {
		return err
	}

	return nil
}

// UpdateUserAvatar update user avatar path
func (d *Dao) UpdateUserAvatar(userID uint64, pathName string) error {
	oldUser, err := d.GetUserByID(userID)
	if err != nil {
		return err
	}

	oldUser.Avatar = pathName

	err = oldUser.Update(d.db)
	if err != nil {
		return err
	}

	return nil
}

// ListUser list usser with pagination
func (d *Dao) ListUser(username, role string, number, page, status int) ([]*model.User, int64, error) {
	if number == 0 {
		number = constvar.DefaultLimit
	}

	where := "`deleted_time` is null"
	whereArgs := []interface{}{}
	if username != "" {
		where += " AND `nickname` LIKE ?"
		whereArgs = append(whereArgs, "%"+username+"%")
	}

	if role != "" {
		where += " AND `role` = ?"
		whereArgs = append(whereArgs, role)
	}

	if status != 0 {
		where += " AND `status` = ?"
		whereArgs = append(whereArgs, status)
	}

	user := &model.User{}
	count, err := user.Count(d.db, where, whereArgs)
	if err != nil {
		return nil, count, err
	}

	offset := (page - 1) * number
	users, err := user.List(d.db, where, whereArgs, offset, number)
	if err != nil {
		return nil, count, err
	}

	return users, count, nil
}

// DeleteUser delete user by ID
func (d *Dao) DeleteUser(userID uint64) error {
	user := &model.User{
		Model: model.Model{ID: userID},
	}
	err := user.Delete(d.db)
	return err
}
