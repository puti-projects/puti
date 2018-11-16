package service

import (
	"strconv"
	"sync"

	"puti/config"
	"puti/model"
	"puti/util"
)

// ArticleInfo is article info for article list
type ArticleInfo struct {
	ID           uint64 `json:"id"`
	UserID       uint64 `json:"userId"`
	Title        string `json:"title"`
	Status       string `json:"status"`
	PostDate     string `json:"post_date"`
	CommentCount uint64 `json:"comment_count"`
	ViewCount    uint64 `json:"view_count"`
}

// ArticleList article list
type ArticleList struct {
	Lock  *sync.Mutex
	IDMap map[uint64]*ArticleInfo
}

// ArticleDetail struct for article detail info
type ArticleDetail struct {
	ID              uint64                 `json:"id"`
	Title           string                 `json:"title"`
	ContentMarkdown string                 `json:"content_markdown"`
	Status          string                 `json:"status"`
	CommentStatus   uint64                 `json:"comment_status"`
	IfTop           uint64                 `json:"if_top"`
	GUID            string                 `json:"guid"`
	CoverPicture    string                 `json:"cover_picture"`
	PostDate        string                 `json:"post_date"`
	MetaData        map[string]interface{} `json:"meta_date"`
	Category        []uint64               `json:"category"`
	Tag             []uint64               `json:"tag"`
}

// ListArticle article list
func ListArticle(title string, page, number int, sort, status string) ([]*ArticleInfo, uint64, error) {
	infos := make([]*ArticleInfo, 0)
	articles, count, err := model.ListArticle(title, page, number, sort, status)
	if err != nil {
		return nil, count, err
	}

	ids := []uint64{}
	for _, article := range articles {
		ids = append(ids, article.ID)
	}

	wg := sync.WaitGroup{}
	articleList := ArticleList{
		Lock:  new(sync.Mutex),
		IDMap: make(map[uint64]*ArticleInfo, len(articles)),
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
			articleList.IDMap[u.ID] = &ArticleInfo{
				ID:           u.ID,
				UserID:       u.UserID,
				Title:        u.Title,
				Status:       u.Status,
				PostDate:     u.PostDate.In(config.TimeLoc()).Format("2006-01-02 15:04"),
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

// GetArticleDetail get article detail by id
func GetArticleDetail(articleID string) (*ArticleDetail, error) {
	ID, _ := strconv.Atoi(articleID)
	uID := uint64(ID)

	// get article info
	a, err := model.GetArticle(uID)
	if err != nil {
		return nil, err
	}

	// get extra data of article
	am, err := model.GetArticleMetaData(uID)
	if err != nil {
		return nil, err
	}

	ArticleDetail := &ArticleDetail{
		ID:              a.ID,
		Title:           a.Title,
		ContentMarkdown: a.ContentMarkdown,
		Status:          a.Status,
		CommentStatus:   a.CommentStatus,
		IfTop:           a.IfTop,
		GUID:            a.GUID,
		CoverPicture:    a.CoverPicture,
		PostDate:        util.GetFormatTime(&a.PostDate, "2006-01-02 15:04:05"),
		MetaData:        make(map[string]interface{}),
	}

	for _, meta := range am {
		ArticleDetail.MetaData[meta.MetaKey] = meta.MetaValue
	}

	articleTaxonomy, err := GetArticleTaxonomy(uID)
	category, categoryOk := articleTaxonomy["category"]
	if categoryOk {
		ArticleDetail.Category = category
	}

	tag, categoryOk := articleTaxonomy["tag"]
	if categoryOk {
		ArticleDetail.Tag = tag
	}

	return ArticleDetail, nil
}
