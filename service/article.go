package service

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"puti/config"
	"puti/model"
	"puti/utils"
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
		PostDate:        utils.GetFormatTime(&a.PostDate, "2006-01-02 15:04:05"),
		MetaData:        make(map[string]interface{}),
		Category:        make([]uint64, 0),
		Tag:             make([]uint64, 0),
	}

	for _, meta := range am {
		ArticleDetail.MetaData[meta.MetaKey] = meta.MetaValue
	}

	articleTaxonomy, err := GetArticleTaxonomy(uID)
	category, categoryOk := articleTaxonomy["category"]
	if categoryOk {
		ArticleDetail.Category = category
	}

	tag, tagOk := articleTaxonomy["tag"]
	if tagOk {
		ArticleDetail.Tag = tag
	}

	return ArticleDetail, nil
}

// UpdateArticle update article info
// In this version, article meta data just update description, it should be more than one choise.TODO
func UpdateArticle(article *model.ArticleModel, description string, category []uint64, tag []uint64) (err error) {
	tx := model.DB.Local.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return err
	}

	// udapte article
	oldArticle, err := model.GetArticle(article.ID)
	if err != nil {
		return err
	}
	oldArticle.Title = article.Title
	oldArticle.ContentMarkdown = article.ContentMarkdown
	oldArticle.ContetnHTML = article.ContetnHTML
	oldArticle.Status = article.Status
	oldArticle.CommentStatus = article.CommentStatus
	oldArticle.IfTop = article.IfTop
	oldArticle.CoverPicture = article.CoverPicture
	oldArticle.PostDate = article.PostDate
	if err = tx.Model(&model.ArticleModel{}).Save(oldArticle).Error; err != nil {
		tx.Rollback()
		return err
	}

	// update article meta data
	oldArticleMeta, err := model.GetOneArticleMetaData(article.ID, "description")
	if oldArticleMeta.MetaValue != description {
		oldArticleMeta.MetaValue = description
		if err = tx.Model(&model.ArticleMetaModel{}).Save(oldArticleMeta).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// get old and new taxonomy
	articleTaxonomy, err := model.GetArticleTaxonomy(article.ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	var oldTaxonomy []uint64
	for _, item := range articleTaxonomy {
		oldTaxonomy = append(oldTaxonomy, item.TermID)
	}
	newTaxonomy := append(category, tag...)

	// delete all old relationship
	dRelation := tx.Where("object_id = ?", article.ID).Delete(model.TermRelationshipsModel{})
	if err := dRelation.Error; err != nil {
		tx.Rollback()
		return err
	}

	// insert all new relationship
	valueStrings := make([]string, 0, len(newTaxonomy))
	valueArgs := make([]interface{}, 0, len(newTaxonomy))
	for _, item := range newTaxonomy {
		termTaxonomy, _ := model.GetTermTaxonomy(item, "")
		valueStrings = append(valueStrings, "(?, ?, ?)")
		valueArgs = append(valueArgs, article.ID)      // object_id
		valueArgs = append(valueArgs, termTaxonomy.ID) // term_taxonomy_id
		valueArgs = append(valueArgs, 0)               // term_order
	}
	tb := &model.TermRelationshipsModel{}
	stmt := fmt.Sprintf("INSERT INTO %s (object_id, term_taxonomy_id, term_order) VALUES %s", tb.TableName(), strings.Join(valueStrings, ","))
	if err := tx.Exec(stmt, valueArgs...).Error; err != nil {
		tx.Rollback()
		return err
	}

	// calculate taxonomy diff
	deleteTaxonomy := calSliceDiff(oldTaxonomy, newTaxonomy)
	insertTaxonomy := calSliceDiff(newTaxonomy, oldTaxonomy)
	// update count
	if err := UpdateTaxonomyCountByArticleChange(tx, insertTaxonomy, 1); err != nil {
		tx.Rollback()
		return err
	}
	if err := UpdateTaxonomyCountByArticleChange(tx, deleteTaxonomy, -1); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func calSliceDiff(slice1, slice2 []uint64) (diffslice []uint64) {
	for _, v := range slice1 {
		inSlice2 := false
		for _, vv := range slice2 {
			if vv == v {
				inSlice2 = true
			}
		}

		if inSlice2 == false {
			diffslice = append(diffslice, v)
		}
	}
	return
}
