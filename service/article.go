package service

import (
	"puti/model"
	"sync"
)

// ListArticle checkout
func ListArticle(title string, page, number int) ([]*model.ArticleInfo, uint64, error) {
	infos := make([]*model.ArticleInfo, 0)
	articles, count, err := model.ListArticle(title, page, number)
	if err != nil {
		return nil, count, err
	}

	ids := []uint64{}
	for _, article := range articles {
		ids = append(ids, article.ID)
	}

	wg := sync.WaitGroup{}
	articleList := model.ArticleList{
		Lock:  new(sync.Mutex),
		IDMap: make(map[uint64]*model.ArticleInfo, len(articles)),
	}

	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	// Improve query efficiency in parallel
	for _, u := range articles {
		wg.Add(1)
		go func(u *model.ArticleModel) {
			defer wg.Done()

			articleList.Lock.Lock()
			defer articleList.Lock.Unlock()
			articleList.IDMap[u.ID] = &model.ArticleInfo{
				ID:           u.ID,
				UserID:       u.UserID,
				Title:        u.Title,
				Status:       u.Status,
				PostDate:     u.PostDate.Format("2006-01-02 15:04:05"),
				CommentCount: u.CommentCount,
				ViewCount:    u.ViewCount,
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
		infos = append(infos, articleList.IDMap[id])
	}

	return infos, count, nil
}
