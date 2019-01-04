package service

import (
	"strings"
	"sync"

	"github.com/puti-projects/puti/internal/common/model"
	"github.com/puti-projects/puti/internal/common/utils"
)

// UserInfo is the user struct for user list
type UserInfo struct {
	ID             uint64 `json:"id"`
	Accout         string `json:"account"`
	Nickname       string `json:"nickname"`
	Email          string `json:"email"`
	Avatar         string `json:"avatar"`
	Roles          string `json:"roles"`
	Status         int    `json:"status"`
	Website        string `json:"website"`
	RegisteredTime string `json:"registered_time"`
	DeletedTime    string `json:"deleted_time"`
}

// UserList user list
type UserList struct {
	Lock  *sync.Mutex
	IDMap map[uint64]*UserInfo
}

// GetUser gets userInfo by username(account)
func GetUser(username string) (*UserInfo, error) {
	u, err := model.GetUser(username)
	if err != nil {
		return nil, err
	}

	userInfo := &UserInfo{
		ID:             u.ID,
		Accout:         u.Username,
		Nickname:       u.Nickname,
		Email:          u.Email,
		Avatar:         u.Avatar,
		Roles:          u.Roles,
		Status:         u.Status,
		Website:        u.PageURL,
		RegisteredTime: utils.GetFormatTime(&u.CreatedAt, "2006-01-02 15:04:05"),
		DeletedTime:    utils.GetFormatTime(u.DeletedAt, "2006-01-02 15:04:05"),
	}

	return userInfo, nil
}

// ListUser show the user list in page
func ListUser(username, role string, number, page, status int) ([]*UserInfo, uint64, error) {
	infos := make([]*UserInfo, 0)
	users, count, err := model.ListUser(username, role, number, page, status)
	if err != nil {
		return nil, count, err
	}

	ids := []uint64{}
	for _, user := range users {
		ids = append(ids, user.ID)
	}

	wg := sync.WaitGroup{}
	userList := UserList{
		Lock:  new(sync.Mutex),
		IDMap: make(map[uint64]*UserInfo, len(users)),
	}

	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	// Improve query efficiency in parallel
	for _, u := range users {
		wg.Add(1)
		go func(u *model.UserModel) {
			defer wg.Done()

			userList.Lock.Lock()
			defer userList.Lock.Unlock()
			userList.IDMap[u.ID] = &UserInfo{
				ID:             u.ID,
				Accout:         u.Username,
				Nickname:       u.Nickname,
				Email:          u.Email,
				Status:         u.Status,
				Roles:          u.Roles,
				RegisteredTime: utils.GetFormatTime(&u.CreatedAt, "2006-01-02 15:04:05"),
				DeletedTime:    utils.GetFormatTime(u.DeletedAt, "2006-01-02 15:04:05")}
		}(u)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChan:
		return nil, count, err
	}

	for _, id := range ids {
		infos = append(infos, userList.IDMap[id])
	}

	return infos, count, nil
}

// UpdateUser updates userinfo by id
func UpdateUser(user *model.UserModel) (err error) {
	// Get old user info
	oldUser, err := model.GetUserByID(user.ID)
	if err != nil {
		return err
	}

	// Set new values
	user.Nickname = strings.TrimSpace(user.Nickname)
	if user.Nickname == "" {
		user.Nickname = oldUser.Username // Set nickname to account if nickname is empty
	}
	oldUser.Nickname = user.Nickname

	oldUser.Email = strings.TrimSpace(user.Email)
	oldUser.PageURL = strings.TrimSpace(user.PageURL)
	oldUser.Roles = user.Roles

	if err = oldUser.Update(); err != nil {
		return err
	}

	return nil
}

// UpdateUserStatus updates user status by id
func UpdateUserStatus(user *model.UserModel) (err error) {
	// Get old user info
	oldUser, err := model.GetUserByID(user.ID)
	if err != nil {
		return err
	}

	// Set new status values
	oldUser.Status = user.Status

	if err = oldUser.Update(); err != nil {
		return err
	}

	return nil
}

// UpdateUserPassword just reset user's password
func UpdateUserPassword(user *model.UserModel) (err error) {
	// Get old user info
	oldUser, err := model.GetUserByID(user.ID)
	if err != nil {
		return err
	}

	// Set new password
	oldUser.Password = user.Password

	if err = oldUser.Update(); err != nil {
		return err
	}

	return nil
}

// UpdateUserAvatar save the new avatar url
func UpdateUserAvatar(user *model.UserModel) (err error) {
	// Get old user info
	oldUser, err := model.GetUserByID(user.ID)
	if err != nil {
		return err
	}

	// Set new password
	oldUser.Avatar = user.Avatar

	if err = oldUser.Update(); err != nil {
		return err
	}

	return nil
}
