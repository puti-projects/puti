package service

import (
	"sync"

	"gingob/model"
)

// ListUser show the user list in page
func ListUser(username string, offset, limit int) ([]*model.UserInfo, uint64, error) {
	infos := make([]*model.UserInfo, 0)
	users, count, err := model.ListUser(username, offset, limit)
	if err != nil {
		return nil, count, err
	}

	ids := []uint64{}
	for _, user := range users {
		ids = append(ids, user.ID)
	}

	wg := sync.WaitGroup{}
	userList := model.UserList{
		Lock:  new(sync.Mutex),
		IDMap: make(map[uint64]*model.UserInfo, len(users)),
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
			userList.IDMap[u.ID] = &model.UserInfo{
				ID:             u.ID,
				Accout:         u.Username,
				Nickname:       u.Nickname,
				Email:          u.Email,
				RegisteredTime: u.RegisteredTime.Format("2006-01-02 15:04:05"),
				Status:         u.Status,
				Role:           u.Role,
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
