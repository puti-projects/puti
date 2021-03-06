package service

import (
	"mime/multipart"
	"strconv"
	"sync"

	"github.com/puti-projects/puti/internal/model"
	"github.com/puti-projects/puti/internal/pkg/config"
	"github.com/puti-projects/puti/internal/pkg/errno"
	"github.com/puti-projects/puti/internal/pkg/token"
	"github.com/puti-projects/puti/internal/utils"

	"github.com/gin-gonic/gin"
)

// UserListRequest is the user list request struct
type UserListRequest struct {
	Username string `form:"username"`
	Number   int    `form:"number"`
	Page     int    `form:"page"`
	Status   int    `form:"status"`
	Role     string `form:"role"`
}

// UserListResponse is the use list response struct
type UserListResponse struct {
	TotalCount int64       `json:"totalCount"`
	UserList   []*UserInfo `json:"userList"`
}

// UserInfo is the user struct for user list
type UserInfo struct {
	ID             uint64 `json:"id"`
	Account        string `json:"account"`
	Nickname       string `json:"nickname"`
	Email          string `json:"email"`
	Avatar         string `json:"avatar"`
	Roles          string `json:"roles"`
	Status         int    `json:"status"`
	Website        string `json:"website"`
	RegisteredTime string `json:"registered_time"`
	DeletedTime    string `json:"deleted_time"`
}

// UserCreateRequest is the create user request params struct
type UserCreateRequest struct {
	Account       string `form:"account"`
	Nickname      string `form:"nickname"`
	Email         string `form:"email"`
	Role          string `form:"role"`
	Password      string `form:"password"`
	PasswordAgain string `form:"passwordAgain"`
	Website       string `form:"website"`
}

// UserCreateResponse is the create user request's response struct
type UserCreateResponse struct {
	Account  string
	Nickname string
}

// GetUser get userInfo by username(account)
func (svc Service) GetUser(username string) (*UserInfo, error) {
	u, err := svc.dao.GetUser(username)
	if err != nil {
		return nil, err
	}
	userInfo := &UserInfo{
		ID:             u.ID,
		Account:        u.Username,
		Nickname:       u.Nickname,
		Email:          u.Email,
		Avatar:         u.Avatar,
		Roles:          u.Roles,
		Status:         u.Status,
		Website:        u.PageURL,
		RegisteredTime: utils.GetFormatTime(&u.CreatedAt, "2006-01-02 15:04:05"),
		DeletedTime:    utils.GetFormatDeletedAtTime(&u.DeletedAt, "2006-01-02 15:04:05"),
	}

	return userInfo, nil
}

// GetUserByToken get userInfo by token(JWT)
func (svc Service) GetUserByToken(t string) (*UserInfo, error) {
	userContext, err := token.ParseToken(t)
	if err != nil {
		return nil, err
	}

	userInfo, err := svc.GetUser(userContext.Username)
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}

// UserList user list handle struct
type UserList struct {
	Lock  *sync.Mutex
	IDMap map[uint64]*UserInfo
}

// ListUser show the user list in page
func (svc Service) ListUser(username, role string, number, page, status int) ([]*UserInfo, int64, error) {
	// get user list
	users, count, err := svc.dao.ListUser(username, role, number, page, status)
	if err != nil {
		return nil, count, err
	}

	infos := make([]*UserInfo, 0)
	ids := []uint64{}
	for _, user := range users {
		ids = append(ids, user.ID)
	}

	wg := sync.WaitGroup{}
	userList := UserList{
		Lock:  new(sync.Mutex),
		IDMap: make(map[uint64]*UserInfo, len(users)),
	}

	finished := make(chan bool, 1)

	// Improve query efficiency in parallel
	for _, u := range users {
		wg.Add(1)
		go func(u *model.User) {
			defer wg.Done()

			userList.Lock.Lock()
			defer userList.Lock.Unlock()
			userList.IDMap[u.ID] = &UserInfo{
				ID:             u.ID,
				Account:        u.Username,
				Nickname:       u.Nickname,
				Email:          u.Email,
				Status:         u.Status,
				Roles:          u.Roles,
				RegisteredTime: utils.GetFormatTime(&u.CreatedAt, "2006-01-02 15:04:05"),
				DeletedTime:    utils.GetFormatDeletedAtTime(&u.DeletedAt, "2006-01-02 15:04:05")}
		}(u)
	}

	go func() {
		// wait for finish
		wg.Wait()
		// close finished channel when finished
		close(finished)
	}()

	<-finished

	for _, id := range ids {
		infos = append(infos, userList.IDMap[id])
	}

	return infos, count, nil
}

// CreateUser create a new user
func (svc Service) CreateUser(u *UserCreateRequest) (string, string, error) {
	if "" == u.Nickname {
		u.Nickname = u.Account
	}

	user := &model.User{
		Username: u.Account,
		Password: u.Password,
		Nickname: u.Nickname,
		Email:    u.Email,
		PageURL:  u.Website,
		Status:   1,
		Roles:    u.Role,
	}

	if err := svc.dao.CreateUser(user); err != nil {
		return "", "", err
	}

	return user.Username, user.Nickname, nil
}

// UserUpdateRequest is the update user request params struct
type UserUpdateRequest struct {
	ID       uint64 `json:"id"`
	Nickname string `json:"nickname"`
	Email    string `json:"email" binding:"required"`
	Role     string `json:"role" binding:"required"`
	Website  string `json:"website"`
}

// UserUpdateStatusRequest only use for update user status
type UserUpdateStatusRequest struct {
	ID     uint64 `json:"id"`
	Status int    `json:"status" binding:"required"`
}

// UserUpdatePasswordRequest user for reset user's password during the profile page
type UserUpdatePasswordRequest struct {
	ID            uint64 `json:"id"`
	Password      string `json:"password" binding:"required"`
	PasswordAgain string `json:"passwordAgain" binding:"required"`
}

// UpdateUser update user info by id
func (svc Service) UpdateUser(u *UserUpdateRequest, userID int) error {
	err := svc.dao.UpdateUser(uint64(userID), u.Nickname, u.Email, u.Website, u.Role)
	if err != nil {
		return errno.New(errno.ErrDatabase, err)
	}

	return nil
}

// UpdateUserStatus update user status
func (svc Service) UpdateUserStatus(u *UserUpdateStatusRequest, userID int) error {
	err := svc.dao.UpdateUserStatus(uint64(userID), u.Status)
	if err != nil {
		return errno.New(errno.ErrDatabase, err)
	}

	return nil
}

// UpdateUserPassword just reset user's password
func (svc Service) UpdateUserPassword(u *UserUpdatePasswordRequest, userID int) error {
	err := svc.dao.UpdateUserPassword(uint64(userID), u.Password)
	if err != nil {
		return errno.New(errno.ErrDatabase, err)
	}

	return nil
}

// UpdateUserAvatar update user avatar
func (svc Service) UpdateUserAvatar(c *gin.Context, userID string, file *multipart.FileHeader) error {
	fileExt := utils.GetFileExt(file)
	newFileName := "user_" + userID + fileExt

	// Upload the file to specific dst.
	pathName := config.UploadUserAvatarPath + newFileName
	dst := "." + pathName
	if err := c.SaveUploadedFile(file, dst); err != nil {
		return errno.New(errno.ErrSaveAvatar, err)
	}

	// update user info for avatar
	ID, err := strconv.Atoi(userID)
	if err != nil {
		return errno.New(errno.ErrSaveAvatar, err)
	}

	err = svc.dao.UpdateUserAvatar(uint64(ID), pathName)
	if err != nil {
		return errno.New(errno.ErrDatabase, err)
	}

	return nil
}

// DeleteUser delete user by ID
func (svc Service) DeleteUser(userID int) error {
	err := svc.dao.DeleteUser(uint64(userID))
	if err != nil {
		return errno.New(errno.ErrDatabase, err)
	}

	return nil
}
