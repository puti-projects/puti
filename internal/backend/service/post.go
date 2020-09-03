package service

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/puti-projects/puti/internal/model"
	"github.com/puti-projects/puti/internal/pkg/db"
	"github.com/puti-projects/puti/internal/utils"

	"github.com/jinzhu/gorm"
)

// PostInfo is post info for post list
type PostInfo struct {
	ID           uint64 `json:"id"`
	UserID       uint64 `json:"userId"`
	Title        string `json:"title"`
	Status       string `json:"status"`
	PostDate     string `json:"post_date"`
	CommentCount uint64 `json:"comment_count"`
	ViewCount    uint64 `json:"view_count"`
}

// PostList post list
type PostList struct {
	Lock  *sync.Mutex
	IDMap map[uint64]*PostInfo
}

// ArticleDetail struct for article info detail
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
	Subject         []uint64               `json:"subject"`
}

// PageDetail struct for page info detail
type PageDetail struct {
	ID              uint64                 `json:"id"`
	Title           string                 `json:"title"`
	ContentMarkdown string                 `json:"content_markdown"`
	Slug            string                 `json:"slug"`
	ParentID        uint64                 `json:"parent_id"`
	Status          string                 `json:"status"`
	CommentStatus   uint64                 `json:"comment_status"`
	GUID            string                 `json:"guid"`
	CoverPicture    string                 `json:"cover_picture"`
	PostDate        string                 `json:"post_date"`
	MetaData        map[string]interface{} `json:"meta_date"`
}

// ListPost post list
func ListPost(postType, title string, page, number int, sort, status string) ([]*PostInfo, uint64, error) {
	infos := make([]*PostInfo, 0)
	posts, count, err := model.ListPost(postType, title, page, number, sort, status)
	if err != nil {
		return nil, count, err
	}

	ids := []uint64{}
	for _, post := range posts {
		ids = append(ids, post.ID)
	}

	wg := sync.WaitGroup{}
	postList := PostList{
		Lock:  new(sync.Mutex),
		IDMap: make(map[uint64]*PostInfo, len(posts)),
	}

	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	// Improve query efficiency in parallel
	for _, u := range posts {
		wg.Add(1)
		go func(u *model.PostModel) {
			defer wg.Done()

			postList.Lock.Lock()
			defer postList.Lock.Unlock()
			postList.IDMap[u.ID] = &PostInfo{
				ID:           u.ID,
				UserID:       u.UserID,
				Title:        u.Title,
				Status:       u.Status,
				PostDate:     utils.GetFormatNullTime(&u.PostDate, "2006-01-02 15:04"),
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
		infos = append(infos, postList.IDMap[id])
	}

	return infos, count, nil
}

// GetArticleDetail get article detail by id
func GetArticleDetail(articleID string) (*ArticleDetail, error) {
	ID, _ := strconv.Atoi(articleID)
	uID := uint64(ID)

	// get article info
	a, err := model.GetPost(uID)
	if err != nil {
		return nil, err
	}

	// get extra data of article
	am, err := model.GetPostMetaData(uID)
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
		PostDate:        utils.GetFormatNullTime(&a.PostDate, "2006-01-02 15:04:05"),
		MetaData:        make(map[string]interface{}),
		Category:        make([]uint64, 0),
		Tag:             make([]uint64, 0),
		Subject:         make([]uint64, 0),
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

	subjectTaxonomy, err := GetArticleSubejct(uID)
	if err != nil {
		return nil, err
	}
	ArticleDetail.Subject = subjectTaxonomy

	return ArticleDetail, nil
}

// GetPageDetail get page detail by id
func GetPageDetail(pageID string) (*PageDetail, error) {
	ID, _ := strconv.Atoi(pageID)
	uID := uint64(ID)

	// get page info
	p, err := model.GetPost(uID)
	if err != nil {
		return nil, err
	}

	// get extra data of page
	pm, err := model.GetPostMetaData(uID)
	if err != nil {
		return nil, err
	}

	pageDetail := &PageDetail{
		ID:              p.ID,
		Title:           p.Title,
		ContentMarkdown: p.ContentMarkdown,
		Slug:            p.Slug,
		ParentID:        p.ParentID,
		Status:          p.Status,
		CommentStatus:   p.CommentStatus,
		GUID:            p.GUID,
		CoverPicture:    p.CoverPicture,
		PostDate:        utils.GetFormatNullTime(&p.PostDate, "2006-01-02 15:04:05"),
		MetaData:        make(map[string]interface{}),
	}

	for _, meta := range pm {
		pageDetail.MetaData[meta.MetaKey] = meta.MetaValue
	}

	return pageDetail, nil
}

// UpdateArticle update article info
// In this version, article meta data just update description, it should be more than one choise.TODO
func UpdateArticle(article *model.PostModel, description string, category []uint64, tag []uint64, subject []uint64) (err error) {
	tx := db.DBEngine.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return err
	}

	// udapte article
	oldArticle, err := model.GetPost(article.ID)
	if err != nil {
		return err
	}
	oldArticle.Title = article.Title
	oldArticle.ContentMarkdown = article.ContentMarkdown
	oldArticle.ContentHTML = article.ContentHTML
	oldArticle.Status = article.Status
	oldArticle.CommentStatus = article.CommentStatus
	oldArticle.IfTop = article.IfTop
	oldArticle.CoverPicture = article.CoverPicture
	oldArticle.PostDate = article.PostDate
	if oldArticle.PostDate.Valid == false && article.Status == model.PostStatusPublish {
		oldArticle.PostDate = mysql.NullTime{Time: time.Now(), Valid: true}
	}
	if err = tx.Model(&model.PostModel{}).Save(oldArticle).Error; err != nil {
		tx.Rollback()
		return err
	}

	// update article meta data
	oldArticleMeta, err := model.GetOnePostMetaData(article.ID, "description")
	if oldArticleMeta.MetaValue != description {
		oldArticleMeta.MetaValue = description
		if err = tx.Model(&model.PostMetaModel{}).Save(oldArticleMeta).Error; err != nil {
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

	// delete all old taxonomy relationship
	dRelation := tx.Where("object_id = ?", article.ID).Delete(model.TermRelationshipsModel{})
	if err := dRelation.Error; err != nil {
		tx.Rollback()
		return err
	}

	// insert all new relationship
	valueStrings := make([]string, 0, len(newTaxonomy))
	valueArgs := make([]interface{}, 0, len(newTaxonomy)*3)
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

	// get old subject
	articleSubject, err := model.GetArticleSubject(article.ID)
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		tx.Rollback()
		return err
	}
	var oldSubject []uint64
	for _, item := range articleSubject {
		oldSubject = append(oldSubject, item.SubjectID)
	}

	// delete all old subject relationship
	deleteSubjectRelation := tx.Where("`object_id` = ?", article.ID).Delete(model.SubjectRelationshipsModel{})
	if err := deleteSubjectRelation.Error; err != nil {
		tx.Rollback()
		return err
	}

	// insert all new relationship
	if subjectLen := len(subject); subjectLen != 0 {
		subjectValueStrings := make([]string, 0, subjectLen)
		subjectValueArgs := make([]interface{}, 0, subjectLen*3)
		for _, subject := range subject {
			subjectValueStrings = append(subjectValueStrings, "(?, ?, ?)")
			subjectValueArgs = append(subjectValueArgs, article.ID) // object_id
			subjectValueArgs = append(subjectValueArgs, subject)    // subject_id
			subjectValueArgs = append(subjectValueArgs, 0)          // order_num
		}
		sr := &model.SubjectRelationshipsModel{}
		sqlr := fmt.Sprintf("INSERT INTO %s (object_id, subject_id, order_num) VALUES %s", sr.TableName(), strings.Join(subjectValueStrings, ","))
		if err := tx.Exec(sqlr, subjectValueArgs...).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// calculate subject diff
	deleteSubject := calSliceDiff(oldSubject, subject)
	insertSubject := calSliceDiff(subject, oldSubject)

	// update subject's info
	if err := UpdateSubjectInfoByArticleChange(tx, insertSubject, 1, true); err != nil {
		tx.Rollback()
		return err
	}
	if err := UpdateSubjectInfoByArticleChange(tx, deleteSubject, -1, true); err != nil {
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

// UpdatePage udpate page info
func UpdatePage(page *model.PostModel, description, pageTemplate string) (err error) {
	tx := db.DBEngine.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return err
	}

	// udapte page
	oldPage, err := model.GetPost(page.ID)
	if err != nil {
		return err
	}
	oldPage.Title = page.Title
	oldPage.ContentMarkdown = page.ContentMarkdown
	oldPage.ContentHTML = page.ContentHTML
	oldPage.Status = page.Status
	oldPage.CommentStatus = page.CommentStatus
	oldPage.CoverPicture = page.CoverPicture
	oldPage.PostDate = page.PostDate
	if oldPage.Slug != page.Slug {
		oldPage.Slug = page.Slug
		oldPage.GUID = fmt.Sprintf("/%s", page.Slug)
	}
	oldPage.ParentID = page.ParentID
	if err = tx.Model(&model.PostModel{}).Save(oldPage).Error; err != nil {
		tx.Rollback()
		return err
	}

	// update article meta data
	// update description
	metaDescription, err := model.GetOnePostMetaData(page.ID, "description")
	if metaDescription.MetaValue != description {
		metaDescription.MetaValue = description
		if err = tx.Model(&model.PostMetaModel{}).Save(metaDescription).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	// update page_template
	metaPageTemplate, err := model.GetOnePostMetaData(page.ID, "page_template")
	if metaPageTemplate.MetaValue != pageTemplate {
		metaPageTemplate.MetaValue = pageTemplate
		if err = tx.Model(&model.PostMetaModel{}).Save(metaPageTemplate).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

// DeletePost delete post by soft delete
// meta data was reserved
func DeletePost(postType string, articleID uint64) error {
	tx := db.DBEngine.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if postType == "article" {
		if err := deleteArticleElse(tx, articleID); err != nil {
			tx.Rollback()
			return err
		}
	}

	// delete post
	if err := tx.Where("id = ?", articleID).Delete(model.PostModel{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// deleteArticleElse delete extra data which only article have
func deleteArticleElse(tx *gorm.DB, articleID uint64) error {
	// get article taxonomy by id
	articleTaxonomy, err := GetArticleTaxonomy(articleID)
	if err != nil {
		return err
	}

	// delete article relationship
	dRelation := tx.Where("`object_id` = ?", articleID).Delete(model.TermRelationshipsModel{})
	if err := dRelation.Error; err != nil {
		return err
	}

	// recount and update taxonomy count
	taxonomy := append(articleTaxonomy["category"], articleTaxonomy["tag"]...)
	if len(taxonomy) != 0 {
		// update category count
		err := UpdateTaxonomyCountByArticleChange(tx, taxonomy, -1)
		if err != nil {
			return err
		}
	}

	// get article subject by id
	articleSubject, err := model.GetArticleSubject(articleID)
	if err != nil {
		return err
	}

	// delete article subject
	dsRelation := tx.Where("`object_id` = ?", articleID).Delete(model.SubjectRelationshipsModel{})
	if err := dsRelation.Error; err != nil {
		return err
	}

	// recount and update subject count
	subjectIDs := make([]uint64, 0)
	for _, subject := range articleSubject {
		subjectIDs = append(subjectIDs, subject.SubjectID)
	}
	if len(subjectIDs) != 0 {
		// update subject count
		if err := UpdateSubjectInfoByArticleChange(tx, subjectIDs, -1, false); err != nil {
			return err
		}
	}

	return nil
}

// TrashPost put the post into the trash by "delete" button
// The different between DeleteArticle and TrashPost is that TrashPost just set the status to deleted
func TrashPost(postID uint64) error {
	oldPost, err := model.GetPost(postID)
	if err != nil {
		return err
	}

	oldPost.Status = "deleted"

	if err = db.DBEngine.Model(&model.PostModel{}).Save(oldPost).Error; err != nil {
		return err
	}

	return nil
}

// RestorePost restore the post which had been put to the trash
// this restore action will set the post as a draft status
func RestorePost(articleID uint64) error {
	oldArticle, err := model.GetPost(articleID)
	if err != nil {
		return err
	}

	oldArticle.Status = "draft"

	if err = db.DBEngine.Model(&model.PostModel{}).Save(oldArticle).Error; err != nil {
		return err
	}

	return nil
}
