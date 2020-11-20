package dao

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/puti-projects/puti/internal/model"
	"github.com/puti-projects/puti/internal/pkg/logger"
	"github.com/puti-projects/puti/internal/utils"

	"gorm.io/gorm"
)

// ArticleTaxonomy use for article taxonomy
type ArticleTaxonomy struct {
	TermID   uint64
	Taxonomy string
}

// GetPostByID get post by post ID
func (d *Dao) GetPostByID(postID uint64) (*model.Post, error) {
	post := &model.Post{Model: model.Model{ID: postID}}
	if err := post.GetByID(d.db); err != nil {
		return nil, err
	}

	return post, nil
}

func (d *Dao) GetPostMetaByPostID(postID uint64) ([]*model.PostMeta, error) {
	pm := model.PostMeta{
		PostID: postID,
	}
	postMeta, err := pm.GetAllByPostID(d.db)
	if err != nil {
		return nil, err
	}

	return postMeta, nil
}

// CreatePage create page
func (d *Dao) CreatePage(page *model.Post, meta []*model.PostMeta) (*model.Post, error) {
	tx := d.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	// create page
	if err := page.Create(tx); err != nil {
		tx.Rollback()
		return nil, err
	}

	// create meta data
	if err := tx.Create(&meta).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return page, tx.Commit().Error
}

// CreateArticle create article
func (d *Dao) CreateArticle(article *model.Post, meta []*model.PostMeta, category, tag, subject []uint64) (*model.Post, error) {
	tx := d.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := article.Create(tx); err != nil {
		tx.Rollback()
		return nil, err
	}

	// set and update GUID column
	article.GUID = fmt.Sprintf("/article/%s.html", strconv.FormatUint(uint64(article.ID), 10))
	if err := article.Save(tx); err != nil {
		tx.Rollback()
		return nil, err
	}

	// set metadata
	var allMeta []*model.PostMeta
	for _, m := range meta {
		m.PostID = article.ID
		allMeta = append(allMeta, m)
	}
	if err := tx.Create(&allMeta).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// set category and tag
	valueStrings := make([]string, 0, len(category)+len(tag))
	valueArgs := make([]interface{}, 0, len(category)+len(tag)*3)
	for _, c := range category {
		termTaxonomy := &model.TermTaxonomy{
			TermID:   c,
			Taxonomy: "category",
		}
		_ = termTaxonomy.GetByTermID(tx)
		valueStrings = append(valueStrings, "(?, ?, ?)")
		valueArgs = append(valueArgs, article.ID)      // object_id
		valueArgs = append(valueArgs, termTaxonomy.ID) // term_taxonomy_id
		valueArgs = append(valueArgs, 0)               // term_order
	}
	for _, t := range tag {
		termTaxonomy := &model.TermTaxonomy{
			TermID:   t,
			Taxonomy: "tag",
		}
		_ = termTaxonomy.GetByTermID(tx)
		valueStrings = append(valueStrings, "(?, ?, ?)")
		valueArgs = append(valueArgs, article.ID)      // object_id
		valueArgs = append(valueArgs, termTaxonomy.ID) // term_taxonomy_id
		valueArgs = append(valueArgs, 0)               // term_order
	}
	tr := &model.TermRelationships{}
	stmt := fmt.Sprintf(
		"INSERT INTO %s (object_id, term_taxonomy_id, term_order) VALUES %s",
		tr.TableName(),
		strings.Join(valueStrings, ","),
	)
	if err := tx.Exec(stmt, valueArgs...).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// update taxonomy count
	insertTaxonomy := append(category, tag...) // combine catogory and tag; they are all taxonomy
	if err := updateTaxonomyCountByArticleChange(tx, insertTaxonomy, 1); err != nil {
		tx.Rollback()
		return nil, err
	}

	// if upload subject id
	if subjectLen := len(subject); subjectLen != 0 {
		subjectValueStrings := make([]string, 0, subjectLen)
		subjectValueArgs := make([]interface{}, 0, subjectLen*3)
		for _, sID := range subject {
			if sID != 0 {
				subjectValueStrings = append(subjectValueStrings, "(?, ?, ?)")
				subjectValueArgs = append(subjectValueArgs, article.ID)
				subjectValueArgs = append(subjectValueArgs, sID)
				subjectValueArgs = append(subjectValueArgs, 0)
			}
		}
		sr := &model.SubjectRelationships{}
		sqls := fmt.Sprintf(
			"INSERT INTO %s (object_id, subject_id, order_num) VALUES %s",
			sr.TableName(),
			strings.Join(subjectValueStrings, ","),
		)
		if err := tx.Exec(sqls, subjectValueArgs...).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		// update subject
		if err := updateSubjectInfoByArticleChange(tx, subject, 1, true); err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// commit
	return article, tx.Commit().Error
}

// UpdatePage update page
func (d *Dao) UpdatePage(page *model.Post, description, pageTemplate string) error {
	// ======================================================
	// ================== Transaction Start =================
	// ======================================================
	tx := d.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// ======================================================
	// Udapte page
	// ======================================================
	if err := page.Save(tx); err != nil {
		tx.Rollback()
		return err
	}

	// ======================================================
	// Update page meta data
	// description,  pageTemplate
	// ======================================================
	pm := &model.PostMeta{
		PostID: page.ID,
	}
	pms, err := pm.GetAllByPostID(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	for k, pm := range pms {
		if pm.MetaKey == "description" && pm.MetaValue != description {
			pms[k].MetaValue = description
		}

		if pm.MetaKey == "page_template" && pm.MetaValue != pageTemplate {
			pms[k].MetaValue = pageTemplate
		}
	}
	for _, pm := range pms {
		if err := pm.Save(tx); err != nil {
			tx.Rollback()
			return err
		}
	}

	// ======================================================
	// ============== Transaction end and commit ============
	// ======================================================
	return tx.Commit().Error
}

// UpdateArticle update article
func (d *Dao) UpdateArticle(article *model.Post, description string, category, tag, subject []uint64) error {
	// ======================================================
	// ================== Transaction Start =================
	// ======================================================
	tx := d.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// ======================================================
	// Udapte article
	// ======================================================
	if err := article.Save(tx); err != nil {
		tx.Rollback()
		return err
	}

	// ======================================================
	// Update article meta data description
	// ======================================================
	articleMeta := model.PostMeta{
		PostID:  article.ID,
		MetaKey: "description",
	}
	if err := articleMeta.GetOneByPostID(tx); err != nil {
		tx.Rollback()
		return err
	}
	if articleMeta.MetaValue != description {
		articleMeta.MetaValue = description
		if err := articleMeta.Save(tx); err != nil {
			tx.Rollback()
			return err
		}
	}

	// ======================================================
	// Update taxonomy
	// ======================================================
	// get old and new taxonomy
	articleTaxonomy, err := getArticleTaxonomyByArticleID(tx, article.ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	var oldTaxonomy []uint64 // TermID group
	for _, item := range articleTaxonomy {
		oldTaxonomy = append(oldTaxonomy, item.TermID)
	}
	newTaxonomy := append(category, tag...) // combine category and tag that input

	// delete all old taxonomy relationship
	mr := &model.TermRelationships{}
	if err := mr.DeleteByCondition(tx, "`object_id` = ?", []interface{}{article.ID}); err != nil {
		tx.Rollback()
		return err
	}
	// insert all new relationship
	valueStrings := make([]string, 0, len(newTaxonomy))
	valueArgs := make([]interface{}, 0, len(newTaxonomy)*3)
	for _, item := range newTaxonomy {
		termTaxonomy := &model.TermTaxonomy{
			TermID: item,
		}
		_ = termTaxonomy.GetByTermID(tx)
		valueStrings = append(valueStrings, "(?, ?, ?)")
		valueArgs = append(valueArgs, article.ID)      // object_id
		valueArgs = append(valueArgs, termTaxonomy.ID) // term_taxonomy_id
		valueArgs = append(valueArgs, 0)               // term_order
	}
	stmt := fmt.Sprintf(
		"INSERT INTO %s (object_id, term_taxonomy_id, term_order) VALUES %s",
		mr.TableName(),
		strings.Join(valueStrings, ","),
	)
	if err := tx.Exec(stmt, valueArgs...).Error; err != nil {
		tx.Rollback()
		return err
	}

	// calculate taxonomy diff
	deleteTaxonomy := utils.SliceDiffUint64(oldTaxonomy, newTaxonomy)
	insertTaxonomy := utils.SliceDiffUint64(newTaxonomy, oldTaxonomy)
	// update count
	if err := updateTaxonomyCountByArticleChange(tx, insertTaxonomy, 1); err != nil {
		tx.Rollback()
		return err
	}
	if err := updateTaxonomyCountByArticleChange(tx, deleteTaxonomy, -1); err != nil {
		tx.Rollback()
		return err
	}

	// ======================================================
	// Update subject
	// ======================================================
	// get old subject
	sr := &model.SubjectRelationships{}
	articleSubject, err := sr.GetAllByObjectID(tx, article.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return err
	}
	var oldSubject []uint64
	for _, item := range articleSubject {
		oldSubject = append(oldSubject, item.SubjectID)
	}

	// delete all old subject relationship
	if err := sr.DeleteByCondition(tx, "`object_id` = ?", []interface{}{article.ID}); err != nil {
		tx.Rollback()
		return err
	}
	// insert all new relationship
	if subjectLen := len(subject); subjectLen != 0 {
		subjectValueStrings := make([]string, 0, subjectLen)
		subjectValueArgs := make([]interface{}, 0, subjectLen*3)
		for _, si := range subject {
			subjectValueStrings = append(subjectValueStrings, "(?, ?, ?)")
			subjectValueArgs = append(subjectValueArgs, article.ID) // object_id
			subjectValueArgs = append(subjectValueArgs, si)         // subject_id
			subjectValueArgs = append(subjectValueArgs, 0)          // order_num
		}
		sqlr := fmt.Sprintf(
			"INSERT INTO %s (object_id, subject_id, order_num) VALUES %s",
			sr.TableName(),
			strings.Join(subjectValueStrings, ","),
		)
		if err := tx.Exec(sqlr, subjectValueArgs...).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// calculate subject diff
	deleteSubject := utils.SliceDiffUint64(oldSubject, subject)
	insertSubject := utils.SliceDiffUint64(subject, oldSubject)

	// update subject's info
	if err := updateSubjectInfoByArticleChange(tx, insertSubject, 1, true); err != nil {
		tx.Rollback()
		return err
	}
	if err := updateSubjectInfoByArticleChange(tx, deleteSubject, -1, true); err != nil {
		tx.Rollback()
		return err
	}

	// ======================================================
	// ============== Transaction end and commit ============
	// ======================================================
	return tx.Commit().Error
}

// TrashPost set post status to "deleted"
func (d *Dao) TrashPost(postID uint64) error {
	post, err := d.GetPostByID(postID)
	if err != nil {
		return err
	}

	// set status to deleted
	post.Status = model.PostStatusDeleted
	return post.Save(d.db)
}

// RestorePost set post status to "draft"
func (d *Dao) RestorePost(postID uint64) error {
	post, err := d.GetPostByID(postID)
	if err != nil {
		return err
	}

	post.Status = model.PostStatusDraft
	return post.Save(d.db)
}

// DeletePost delete post
func (d *Dao) DeletePost(postType string, articleID uint64) error {
	tx := d.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	if postType == model.PostTypeArticle {
		// delete some article related data
		if err := d.deleteArticleElse(tx, articleID); err != nil {
			tx.Rollback()
			return err
		}
	}

	// delete post
	post := model.Post{Model: model.Model{ID: articleID}}
	if err := post.Delete(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// deleteArticleElse delete extra data which only article have
func (d *Dao) deleteArticleElse(tx *gorm.DB, articleID uint64) error {
	// get article taxonomy by id
	articleTaxonomy, err := d.GetArticleTaxonomy(tx, articleID)
	if err != nil {
		return err
	}

	// delete article relationship
	tr := &model.TermRelationships{}
	if err := tr.DeleteByCondition(tx, "`object_id` = ?", []interface{}{articleID}); err != nil {
		return err
	}

	// recount and update taxonomy count
	taxonomy := append(articleTaxonomy["category"], articleTaxonomy["tag"]...)
	if len(taxonomy) != 0 {
		// update category count
		err := updateTaxonomyCountByArticleChange(tx, taxonomy, -1)
		if err != nil {
			return err
		}
	}

	// get article subject by id
	sr := &model.SubjectRelationships{}
	articleSubject, err := sr.GetAllByObjectID(tx, articleID)
	if err != nil {
		return err
	}
	// delete article subject
	if err := sr.DeleteByCondition(tx, "`object_id` = ?", []interface{}{articleID}); err != nil {
		return err
	}

	// recount and update subject count
	subjectIDGroup := make([]uint64, 0)
	for _, subject := range articleSubject {
		subjectIDGroup = append(subjectIDGroup, subject.SubjectID)
	}
	if len(subjectIDGroup) != 0 {
		// update subject count
		if err := updateSubjectInfoByArticleChange(tx, subjectIDGroup, -1, false); err != nil {
			return err
		}
	}

	return nil
}

// GetArticleTaxonomyByArticleID get article taxonomy include all type
func getArticleTaxonomyByArticleID(tx *gorm.DB, articleID uint64) ([]*ArticleTaxonomy, error) {
	sql := "SELECT t.term_id, tt.taxonomy FROM pt_term t LEFT JOIN pt_term_taxonomy tt ON tt.term_id = t.term_id LEFT JOIN pt_term_relationships tr ON tr.term_taxonomy_id = tt.term_taxonomy_id WHERE tr.object_id = ?"
	rows, err := tx.Raw(sql, articleID).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	at := make([]*ArticleTaxonomy, 0)
	for rows.Next() {
		articleTaxonomy := &ArticleTaxonomy{}
		if err := tx.ScanRows(rows, &articleTaxonomy); err != nil {
			return nil, err
		}
		at = append(at, articleTaxonomy)
	}

	return at, nil
}

// GetArticleTaxonomy get taxonomy and output by taxonomy type
func (d *Dao) GetArticleTaxonomy(tx *gorm.DB, articleID uint64) (map[string][]uint64, error) {
	if nil == tx {
		tx = d.db
	}

	taxonomy, err := getArticleTaxonomyByArticleID(tx, articleID)
	if err != nil {
		return nil, err
	}

	articleTaxonomy := make(map[string][]uint64)
	for _, item := range taxonomy {
		_, ok := articleTaxonomy[item.Taxonomy]
		if !ok {
			articleTaxonomy[item.Taxonomy] = make([]uint64, 0)
		}
		articleTaxonomy[item.Taxonomy] = append(articleTaxonomy[item.Taxonomy], item.TermID)
	}

	return articleTaxonomy, nil
}

// ListPost returns the posts list in condition
func (d *Dao) ListPost(postType, title string, page, number int, sort, status string) ([]*model.Post, int64, error) {
	// count
	where := "`post_type` = ? AND `parent_id` = ?"
	whereArgs := []interface{}{postType, 0}
	if "" != title {
		where += " AND `title` LIKE ?"
		whereArgs = append(whereArgs, "%"+title+"%")
	}
	if "" != status {
		where += " AND `status`= ?"
		whereArgs = append(whereArgs, status)
	}
	post := &model.Post{}
	count, err := post.Count(d.db, where, whereArgs)
	if err != nil {
		return nil, count, err
	}

	// list
	offset := (page - 1) * number
	var order string
	if "" != sort {
		order = "id " + sort
	} else {
		order = "id DESC"
	}
	posts, err := post.List(d.db, where, whereArgs, offset, number, order)
	if err != nil {
		return nil, count, err
	}

	return posts, count, nil
}

// CheckPageSlugExist check slug name exist
func (d *Dao) CheckPageSlugExist(pageID uint64, slug string) bool {
	post := &model.Post{
		Model: model.Model{ID: pageID},
		Slug:  slug,
	}
	return post.CheckSlug(d.db)
}

// GetTotalView get total views of allpost
func (d *Dao) GetPostTotalView() (totalViews int64) {
	var err error

	post := &model.Post{}
	totalViews, err = post.TotalView(d.db)
	if err != nil {
		logger.Errorf("statistics error: get total view failed, %v", err)
	}

	return
}

// GetTotalArticles get total number of article
func (d *Dao) GetTotalArticles() (totalPost int64) {
	var err error

	post := &model.Post{}
	totalPost, err = post.TotalNumber(d.db, model.PostTypeArticle)
	if err != nil {
		logger.Errorf("statistics error: get total article failed, %v", err)
	}

	return
}
