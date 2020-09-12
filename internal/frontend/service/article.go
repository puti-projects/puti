package service

import (
	"database/sql"
	"fmt"
	"html/template"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/puti-projects/puti/internal/model"
	"github.com/puti-projects/puti/internal/pkg/config"
	"github.com/puti-projects/puti/internal/pkg/db"
	"github.com/puti-projects/puti/internal/pkg/logger"
	optionCache "github.com/puti-projects/puti/internal/pkg/option"
	"github.com/puti-projects/puti/internal/utils"
)

// List post list
type List struct {
	Lock  *sync.Mutex
	IDMap map[uint64]*ShowArticle
}

// GetArticleListByTaxonomy get articles list
func GetArticleListByTaxonomy(currentPage int, taxonomyType, taxonomySlug, keyword string) (termName string, articleResult []*ShowArticle, pagination *utils.Pagination, err error) {
	// get articles data
	pageSize, _ := strconv.Atoi(optionCache.Options.Get("posts_per_page"))
	offset := (currentPage - 1) * pageSize
	var count int64 = 0

	// get term id
	var termTaxonomyID uint64
	getTermTaxonomyID := db.DBEngine.Table("pt_term as t").
		Select("t.`name`, tt.`term_taxonomy_id`").
		Joins("INNER JOIN pt_term_taxonomy as tt ON tt.term_id = t.term_id").
		Where("t.`slug` = ? AND tt.`taxonomy` = ?", taxonomySlug, taxonomyType).
		Row()
	getTermTaxonomyID.Scan(&termName, &termTaxonomyID)

	// get article list
	where := "p.`deleted_time` IS NULL AND p.`post_type` = ? AND p.`parent_id` = ? AND p.`status` = ? AND tr.`term_taxonomy_id` = ?"
	whereArgs := []interface{}{model.PostTypeArticle, 0, model.PostStatusPublish, termTaxonomyID}
	if "" != keyword {
		where += " AND p.`title` LIKE ?"
		whereArgs = append(whereArgs, "%"+keyword+"%")
	}

	articles := []*model.Post{}
	result := db.DBEngine.Table("pt_post AS p").
		Select("p.`id`, p.`title`, p.`if_top`, p.`content_html`, p.`guid`, p.`cover_picture`, p.`comment_count`, p.`view_count`, p.`posted_time`").
		Joins("INNER JOIN pt_term_relationships AS tr ON tr.object_id = p.id").
		Unscoped().
		Where(where, whereArgs...).Count(&count).
		Order("p.`if_top` DESC, p.`posted_time` DESC").
		Offset(offset).Limit(pageSize).
		Find(&articles)

	if err := result.Error; err != nil {
		logger.Errorf("get articles failed. %s", err)
	}

	// get pagination
	pagination = utils.GetPagination(int(count), currentPage, pageSize, 0)

	// handle articles data
	articleResult = make([]*ShowArticle, 0)
	ids := []uint64{}
	for _, article := range articles {
		ids = append(ids, article.ID)
	}

	wg := sync.WaitGroup{}
	articleList := List{
		Lock:  new(sync.Mutex),
		IDMap: make(map[uint64]*ShowArticle, len(articles)),
	}

	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	siteURL := optionCache.Options.Get("site_url")
	for _, a := range articles {
		wg.Add(1)
		go func(a *model.Post) {
			defer wg.Done()

			var ifTop = false
			if a.IfTop == 1 {
				ifTop = true
			}

			articleCategory, articleTag, err := getArticleTaxonomyInfo(a.ID, siteURL)
			if err != nil {
				errChan <- err
			}

			articleList.Lock.Lock()
			defer articleList.Lock.Unlock()
			articleList.IDMap[a.ID] = &ShowArticle{
				ID:           a.ID,
				Title:        a.Title,
				IfTop:        ifTop,
				Abstract:     getArticleAbstract(a.ContentHTML),
				GUID:         a.GUID,
				CoverPicture: a.CoverPicture,
				CommentCount: a.CommentCount,
				ViewCount:    a.ViewCount,
				PostedTime:   utils.GetFormatNullTime(&a.PostDate, "2006-01-02 15:04"),
				Tags:         articleTag,
				Categories:   articleCategory,
			}
		}(a)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChan:
		return "", nil, nil, err
	}

	for _, id := range ids {
		articleResult = append(articleResult, articleList.IDMap[id])
	}

	return
}

// GetArticleList get articles list
func GetArticleList(currentPage int, keyword string) (articleResult []*ShowArticle, pagination *utils.Pagination, err error) {
	// get articles data
	pageSize, _ := strconv.Atoi(optionCache.Options.Get("posts_per_page"))
	offset := (currentPage - 1) * pageSize
	var count int64 = 0

	where := "`post_type` = ? AND `parent_id` = ? AND `status` = ?"
	whereArgs := []interface{}{model.PostTypeArticle, 0, model.PostStatusPublish}
	if "" != keyword {
		where += " AND `title` LIKE ?"
		whereArgs = append(whereArgs, "%"+keyword+"%")
	}

	articles := []*model.Post{}
	result := db.DBEngine.Model(&model.Post{}).
		Select("`id`, `title`, `if_top`, `content_html`, `guid`, `cover_picture`, `comment_count`, `view_count`, `posted_time`").
		Where(where, whereArgs...).Count(&count).
		Order("`if_top` DESC, `posted_time` DESC").
		Offset(offset).Limit(pageSize).
		Find(&articles)

	if err := result.Error; err != nil {
		logger.Errorf("get articles failed. %s", err)
	}

	// get pagination
	pagination = utils.GetPagination(int(count), currentPage, pageSize, 0)

	// handle articles data
	articleResult = make([]*ShowArticle, 0)
	ids := []uint64{}
	for _, article := range articles {
		ids = append(ids, article.ID)
	}

	wg := sync.WaitGroup{}
	articleList := List{
		Lock:  new(sync.Mutex),
		IDMap: make(map[uint64]*ShowArticle, len(articles)),
	}

	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	siteURL := optionCache.Options.Get("site_url")
	for _, a := range articles {
		wg.Add(1)
		go func(a *model.Post) {
			defer wg.Done()

			var ifTop = false
			if a.IfTop == 1 {
				ifTop = true
			}

			articleCategory, articleTag, err := getArticleTaxonomyInfo(a.ID, siteURL)
			if err != nil {
				errChan <- err
			}

			articleList.Lock.Lock()
			defer articleList.Lock.Unlock()
			articleList.IDMap[a.ID] = &ShowArticle{
				ID:           a.ID,
				Title:        a.Title,
				IfTop:        ifTop,
				Abstract:     getArticleAbstract(a.ContentHTML),
				GUID:         a.GUID,
				CoverPicture: a.CoverPicture,
				CommentCount: a.CommentCount,
				ViewCount:    a.ViewCount,
				PostedTime:   utils.GetFormatNullTime(&a.PostDate, "2006-01-02 15:04"),
				Tags:         articleTag,
				Categories:   articleCategory,
			}
		}(a)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChan:
		return nil, nil, err
	}

	for _, id := range ids {
		articleResult = append(articleResult, articleList.IDMap[id])
	}

	return
}

func getArticleTaxonomyInfo(articleID uint64, siteURL string) ([]*ShowCategory, []*ShowTag, error) {
	sql := "SELECT t.`name`, t.`slug`, tt.`taxonomy` FROM pt_term t LEFT JOIN pt_term_taxonomy tt ON tt.`term_id` = t.`term_id` LEFT JOIN pt_term_relationships tr ON tr.`term_taxonomy_id` = tt.`term_taxonomy_id` WHERE tr.`object_id` = ?"
	rows, err := db.DBEngine.Raw(sql, articleID).Rows()
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	articleCategory := make([]*ShowCategory, 0)
	articleTag := make([]*ShowTag, 0)
	var name string
	var slug string
	var taxonomy string
	for rows.Next() {
		rows.Scan(&name, &slug, &taxonomy)
		if taxonomy == "category" {
			articleCategory = append(articleCategory, &ShowCategory{Title: name, URL: siteURL + config.PathCategory + "/" + slug})
		}

		if taxonomy == "tag" {
			articleTag = append(articleTag, &ShowTag{Title: name, URL: siteURL + config.PathTag + "/" + slug})
		}
	}

	return articleCategory, articleTag, nil
}

func getArticleAbstract(content string) string {
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	content = re.ReplaceAllString(content, "\n")

	re, _ = regexp.Compile("\\s{2,}")
	content = re.ReplaceAllString(content, "\n")

	content = strings.TrimSpace(content)

	abstractRune := []rune(content)
	contentLen := len(abstractRune)
	if contentLen <= 200 {
		return content
	}

	abstract := fmt.Sprintf("%s%s", string(abstractRune[:200]), "...")
	return abstract
}

// GetLatestArticlesList get latest article list for widget
func GetLatestArticlesList(getNums int) ([]*ShowWidgetLatestArticles, error) {
	where := "`post_type` = ? AND `parent_id` = ? AND `status` = ?"
	whereArgs := []interface{}{model.PostTypeArticle, 0, model.PostStatusPublish}

	articles := []*ShowWidgetLatestArticles{}
	postModel := &model.Post{}
	result := db.DBEngine.Table(postModel.TableName()).
		Select("`id`, `title`, `guid`, `comment_count`, `view_count`, `posted_time`").
		Where(where, whereArgs...).
		Order("`posted_time` DESC").
		Limit(getNums).
		Find(&articles)

	if err := result.Error; err != nil {
		return nil, err
	}

	return articles, nil
}

// GetArticleDetailByID get article detail by article id
func GetArticleDetailByID(articleID uint64) (*ShowArticleDetail, error) {
	a := &model.Post{}
	err := db.DBEngine.Where("id = ? AND post_type = ? AND parent_id = ? AND status =?", articleID, model.PostTypeArticle, 0, model.PostStatusPublish).First(&a).Error
	if err != nil {
		return nil, err
	}

	siteURL := optionCache.Options.Get("site_url")
	articleCategory, articleTag, err := getArticleTaxonomyInfo(articleID, siteURL)
	if err != nil {
		return nil, err
	}

	articleDetail := &ShowArticleDetail{
		ID:            a.ID,
		Title:         a.Title,
		ContentHTML:   template.HTML(a.ContentHTML),
		CommentStatus: a.CommentStatus,
		GUID:          siteURL + a.GUID,
		CommentCount:  a.CommentCount,
		ViewCount:     a.ViewCount,
		PostedTime:    utils.GetFormatNullTime(&a.PostDate, "2006-01-02 15:04"),
		MetaData:      make(map[string]interface{}),
		Categories:    articleCategory,
		Tags:          articleTag,
	}

	// get extra data of article
	pm := &model.PostMeta{PostID: articleID}
	am, err := pm.GetAllByPostID(db.DBEngine)
	if err != nil {
		return nil, err
	}
	for _, meta := range am {
		articleDetail.MetaData[meta.MetaKey] = meta.MetaValue
	}

	return articleDetail, nil
}

// GetLastArticle get last article title and url
func GetLastArticle(currentArticleID uint64) *ShowLastOrNextArticle {
	var title string
	var url string
	postModel := &model.Post{}
	row := db.DBEngine.Table(postModel.TableName()).
		Where("`id` < ? AND `post_type` = ? AND `parent_id` = ? AND `status` = ? AND `deleted_time` IS NULL", currentArticleID, model.PostTypeArticle, 0, model.PostStatusPublish).
		Select("`title`, `guid`").Order("`id` DESC").Row()
	row.Scan(&title, &url)

	article := &ShowLastOrNextArticle{
		Title: title,
		URL:   url,
	}

	return article
}

// GetNextArticle get next article title and url
func GetNextArticle(currentArticleID uint64) *ShowLastOrNextArticle {
	var title string
	var url string
	postModel := &model.Post{}
	row := db.DBEngine.Table(postModel.TableName()).
		Where("`id` > ? AND `post_type` = ? AND `parent_id` = ? AND `status` = ? AND `deleted_time` IS NULL", currentArticleID, model.PostTypeArticle, 0, model.PostStatusPublish).
		Select("`title`, `guid`").Order("`id` ASC").Row()
	row.Scan(&title, &url)

	article := &ShowLastOrNextArticle{
		Title: title,
		URL:   url,
	}

	return article
}

// GetSubjectArticleList get article list connected to subject
func GetSubjectArticleList(subjectID uint64) ([]*map[string]interface{}, error) {
	var articleList []*map[string]interface{}

	postModel := &model.Post{}
	srModel := &model.SubjectRelationships{}
	rows, err := db.DBEngine.Table(postModel.TableName()+" p").
		Select("p.`id`, p.`title`, p.`guid`, p.`comment_count`, p.`view_count`, p.`posted_time`").
		Joins("INNER JOIN "+srModel.TableName()+" sr ON sr.`object_id` = p.`id`").
		Where("p.`post_type` = ? AND p.`parent_id` = ? AND p.`status` = ? AND sr.`subject_id` = ? AND p.`deleted_time` is null",
			model.PostTypeArticle, 0, model.PostStatusPublish, subjectID).
		Order("p.`posted_time` DESC").
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var id, commentCount, viewCount uint64
		var title, guid string
		var postedTime sql.NullTime
		rows.Scan(&id, &title, &guid, &commentCount, &viewCount, &postedTime)

		item := &map[string]interface{}{
			"ID":           id,
			"Title":        title,
			"GUID":         guid,
			"CommentCount": commentCount,
			"ViewCount":    viewCount,
			"PostedTime":   postedTime,
		}
		articleList = append(articleList, item)
	}

	return articleList, nil
}
