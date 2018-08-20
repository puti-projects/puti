package service

import (
	"sync"

	"puti/config"
	"puti/model"
)

// UserInfo is the user struct for user list
type UserInfo struct {
	ID             uint64 `json:"id"`
	Accout         string `json:"account"`
	Nickname       string `json:"nickname"`
	Email          string `json:"email"`
	Avator         string `json:"avator"`
	RegisteredTime string `json:"registered_time"`
	Roles          string `json:"roles"`
	Status         int    `json:"status"`
	Website        string `json:"website"`
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
		Avator:         u.Avatar,
		RegisteredTime: u.CreatedAt.In(config.TimeLoc()).Format("2006-01-02 15:04:05"),
		Roles:          u.Roles,
		Status:         u.Status,
		Website:        u.PageURL,
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
				RegisteredTime: u.CreatedAt.In(config.TimeLoc()).Format("2006-01-02 15:04:05"),
				Status:         u.Status,
				Roles:          u.Roles,
			}
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
