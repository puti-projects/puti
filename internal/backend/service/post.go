package service

// article and page handle are all in this post module

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/puti-projects/puti/internal/backend/dao"
	"github.com/puti-projects/puti/internal/model"
	"github.com/puti-projects/puti/internal/pkg/errno"
	"github.com/puti-projects/puti/internal/utils"
)

// ArticleCreateRequest struct of article create params
type ArticleCreateRequest struct {
	Status        string   `json:"status"`
	Title         string   `json:"title"`
	Content       string   `json:"content"`
	ContentHTML   string   `json:"content_html"`
	Description   string   `json:"description"`
	CommentStatus uint64   `json:"comment_status"`
	CoverPicture  string   `json:"cover_picture"`
	PostedTime    string   `json:"posted_time"`
	IfTop         uint64   `json:"if_top"`
	Category      []uint64 `json:"category"`
	Tag           []uint64 `json:"tag"`
	Subject       []uint64 `json:"subject"`
}

// ArticleCreateResponse return the new article id and url
type ArticleCreateResponse struct {
	ID   uint64 `json:"id"`
	GUID string `json:"guid"`
}

// PageCreateRequest struct of page create params
type PageCreateRequest struct {
	Status        string `json:"status"`
	Title         string `json:"title"`
	Content       string `json:"content"`
	ContentHTML   string `json:"content_html"`
	Description   string `json:"description"`
	CommentStatus uint64 `json:"comment_status"`
	CoverPicture  string `json:"cover_picture"`
	PostedTime    string `json:"posted_time"`
	Slug          string `json:"slug"`
	PageTemplate  string `json:"page_template"`
	ParentID      uint64 `json:"parent_id"`
}

// PageCreateResponse return the new page id and url
type PageCreateResponse struct {
	ID   uint64 `json:"id"`
	GUID string `json:"guid"`
}

// ArticleUpdateRequest struct for update article
type ArticleUpdateRequest struct {
	ID            uint64   `json:"id"`
	Status        string   `json:"status"`
	Title         string   `json:"title"`
	Content       string   `json:"content"`
	ContentHTML   string   `json:"content_html"`
	Description   string   `json:"description"`
	CommentStatus uint64   `json:"comment_status"`
	CoverPicture  string   `json:"cover_picture"`
	PostedTime    string   `json:"posted_time"`
	IfTop         uint64   `json:"if_top"`
	Category      []uint64 `json:"category"`
	Tag           []uint64 `json:"tag"`
	Subject       []uint64 `json:"subject"`
}

// PageUpdateRequest struct of page update params
type PageUpdateRequest struct {
	ID            uint64 `json:"id"`
	Status        string `json:"status"`
	Title         string `json:"title"`
	Content       string `json:"content"`
	ContentHTML   string `json:"content_html"`
	Description   string `json:"description"`
	CommentStatus uint64 `json:"comment_status"`
	CoverPicture  string `json:"cover_picture"`
	PostedTime    string `json:"posted_time"`
	Slug          string `json:"slug"`
	PageTemplate  string `json:"page_template"`
	ParentID      uint64 `json:"parent_id"`
}

// ArticleListRequest is the article list request struct
type ArticleListRequest struct {
	Title  string `form:"title"`
	Page   int    `form:"page"`
	Number int    `form:"number"`
	Sort   string `form:"sort"`
	Status string `form:"status"`
}

// ArticleListResponse is the article list response struct
type ArticleListResponse struct {
	TotalCount  int64       `json:"totalCount"`
	TotalPage   uint64      `json:"totalPage"`
	ArticleList []*PostInfo `json:"articleList"`
}

// PageListRequest is the page list request struct
type PageListRequest struct {
	Title  string `form:"title"`
	Page   int    `form:"page"`
	Number int    `form:"number"`
	Sort   string `form:"sort"`
	Status string `form:"status"`
}

// PageListResponse is the page list response struct
type PageListResponse struct {
	TotalCount int64       `json:"totalCount"`
	TotalPage  uint64      `json:"totalPage"`
	PageList   []*PostInfo `json:"pageList"`
}

// PostList post list
type PostList struct {
	Lock  *sync.Mutex
	IDMap map[uint64]*PostInfo
}

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

// CreateArticle create article
func CreateArticle(r *ArticleCreateRequest, userID uint64) (*ArticleCreateResponse, error) {
	// post data
	article := &model.Post{
		UserID:          userID,
		PostType:        model.PostTypeArticle,
		Title:           r.Title,
		ContentMarkdown: r.Content,
		ContentHTML:     r.ContentHTML,
		ParentID:        0,
		Status:          r.Status,
		CommentStatus:   r.CommentStatus,
		IfTop:           r.IfTop,
		CoverPicture:    r.CoverPicture,
		CommentCount:    0,
		ViewCount:       0,
	}
	if r.PostedTime == "" && r.Status == model.PostStatusPublish {
		article.PostDate = sql.NullTime{Time: time.Now(), Valid: true}
	} else {
		article.PostDate = utils.StringToNullTime("2006-01-02 15:04:05", r.PostedTime)
	}

	// post meta data
	descriptionMeta := []*model.PostMeta{
		{
			MetaKey:   "description",
			MetaValue: r.Description,
		},
	}

	article, err := dao.Engine.CreateArticle(article, descriptionMeta, r.Category, r.Tag, r.Subject)
	if err != nil {
		return nil, errno.New(errno.ErrDatabase, err)
	}

	rsp := &ArticleCreateResponse{
		ID:   article.ID,
		GUID: article.GUID,
	}
	return rsp, nil
}

// CreatePage create page
func CreatePage(r *PageCreateRequest, userID uint64) (*PageCreateResponse, error) {
	// page data
	page := &model.Post{
		UserID:          userID,
		PostType:        model.PostTypePage,
		Title:           r.Title,
		ContentMarkdown: r.Content,
		ContentHTML:     r.ContentHTML,
		Slug:            r.Slug,
		ParentID:        r.ParentID,
		Status:          r.Status,
		CommentStatus:   r.CommentStatus,
		IfTop:           0,
		GUID:            fmt.Sprintf("/%s", r.Slug),
		CoverPicture:    r.CoverPicture,
		CommentCount:    0,
		ViewCount:       0,
	}
	if r.PostedTime == "" && r.Status == model.PostStatusPublish {
		page.PostDate = sql.NullTime{Time: time.Now(), Valid: true}
	} else {
		page.PostDate = utils.StringToNullTime("2006-01-02 15:04:05", r.PostedTime)
	}

	// set metadata description
	meta := []*model.PostMeta{
		{
			PostID:    page.ID,
			MetaKey:   "description",
			MetaValue: r.Description,
		},
		{
			PostID:    page.ID,
			MetaKey:   "page_template",
			MetaValue: r.PageTemplate,
		},
	}

	page, err := dao.Engine.CreatePage(page, meta)
	if err != nil {
		return nil, errno.New(errno.ErrDatabase, err)
	}

	rsp := &PageCreateResponse{
		ID:   page.ID,
		GUID: page.GUID,
	}
	return rsp, nil
}

// ListArticle article list
func ListArticle(postType string, r *ArticleListRequest) ([]*PostInfo, int64, error) {
	return ListPost(postType, r.Title, r.Page, r.Number, r.Sort, r.Status)
}

// ListPage page list
func ListPage(postType string, r *PageListRequest) ([]*PostInfo, int64, error) {
	return ListPost(postType, r.Title, r.Page, r.Number, r.Sort, r.Status)
}

// ListPost post list
func ListPost(postType, title string, page, number int, sort, status string) ([]*PostInfo, int64, error) {
	infos := make([]*PostInfo, 0)
	posts, count, err := dao.Engine.ListPost(postType, title, page, number, sort, status)
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

	finished := make(chan bool, 1)

	// Improve query efficiency in parallel
	for _, u := range posts {
		wg.Add(1)
		go func(u *model.Post) {
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

	<-finished

	for _, id := range ids {
		infos = append(infos, postList.IDMap[id])
	}

	return infos, count, nil
}

// CheckPageSlugExist check if page slug name exist
func CheckPageSlugExist(pageID uint64, slug string) bool {
	return dao.Engine.CheckPageSlugExist(pageID, slug)
}

// GetArticleDetail get article detail by id
func GetArticleDetail(articleID uint64) (*ArticleDetail, error) {
	// get article
	article, err := dao.Engine.GetPostByID(articleID)
	if err != nil {
		return nil, err
	}

	// get extra data of article
	articleMeta, err := dao.Engine.GetPostMetaByPostID(articleID)
	if err != nil {
		return nil, err
	}

	// main data
	ArticleDetail := &ArticleDetail{
		ID:              article.ID,
		Title:           article.Title,
		ContentMarkdown: article.ContentMarkdown,
		Status:          article.Status,
		CommentStatus:   article.CommentStatus,
		IfTop:           article.IfTop,
		GUID:            article.GUID,
		CoverPicture:    article.CoverPicture,
		PostDate:        utils.GetFormatNullTime(&article.PostDate, "2006-01-02 15:04:05"),
		MetaData:        make(map[string]interface{}),
		Category:        make([]uint64, 0),
		Tag:             make([]uint64, 0),
		Subject:         make([]uint64, 0),
	}
	// meta data
	for _, meta := range articleMeta {
		ArticleDetail.MetaData[meta.MetaKey] = meta.MetaValue
	}
	// taxonomy data
	articleTaxonomy, err := dao.Engine.GetArticleTaxonomy(nil, articleID)
	if err != nil {
		return nil, err
	}
	category, categoryOk := articleTaxonomy["category"]
	if categoryOk {
		ArticleDetail.Category = category
	}
	tag, tagOk := articleTaxonomy["tag"]
	if tagOk {
		ArticleDetail.Tag = tag
	}
	// subject
	articleSubjectGroup, err := GetArticleSubejctID(articleID)
	if err != nil {
		return nil, err
	}
	ArticleDetail.Subject = articleSubjectGroup

	return ArticleDetail, nil
}

// GetPageDetail get page detail by id
func GetPageDetail(pageID uint64) (*PageDetail, error) {
	// get page
	page, err := dao.Engine.GetPostByID(pageID)
	if err != nil {
		return nil, err
	}

	// get extra data of page
	pageMeta, err := dao.Engine.GetPostMetaByPostID(pageID)
	if err != nil {
		return nil, err
	}

	pageDetail := &PageDetail{
		ID:              page.ID,
		Title:           page.Title,
		ContentMarkdown: page.ContentMarkdown,
		Slug:            page.Slug,
		ParentID:        page.ParentID,
		Status:          page.Status,
		CommentStatus:   page.CommentStatus,
		GUID:            page.GUID,
		CoverPicture:    page.CoverPicture,
		PostDate:        utils.GetFormatNullTime(&page.PostDate, "2006-01-02 15:04:05"),
		MetaData:        make(map[string]interface{}),
	}

	for _, meta := range pageMeta {
		pageDetail.MetaData[meta.MetaKey] = meta.MetaValue
	}

	return pageDetail, nil
}

// UpdateArticle update article
// TODO: In this version, article meta data just update description, it should be more than one choise.
func UpdateArticle(a *ArticleUpdateRequest) error {
	// get old article data
	article, err := dao.Engine.GetPostByID(a.ID)
	if err != nil {
		return err
	}

	// reset article data
	article.Title = a.Title
	article.ContentMarkdown = a.Content
	article.ContentHTML = a.ContentHTML
	article.Status = a.Status
	article.CommentStatus = a.CommentStatus
	article.IfTop = a.IfTop
	article.CoverPicture = a.CoverPicture
	article.PostDate = utils.StringToNullTime("2006-01-02 15:04:05", a.PostedTime)
	if article.PostDate.Valid == false && article.Status == model.PostStatusPublish {
		// first publish; save as draft and  publish now in this situation
		article.PostDate = sql.NullTime{Time: time.Now(), Valid: true}
	}

	err = dao.Engine.UpdateArticle(article, a.Description, a.Category, a.Tag, a.Subject)
	if err != nil {
		return errno.New(errno.ErrDatabase, err)
	}

	return nil
}

// UpdatePage update page
func UpdatePage(p *PageUpdateRequest) (err error) {
	// get old page data
	page, err := dao.Engine.GetPostByID(p.ID)
	if err != nil {
		return err
	}

	page.Title = p.Title
	page.ContentMarkdown = p.Content
	page.ContentHTML = p.ContentHTML
	page.Status = p.Status
	page.CommentStatus = p.CommentStatus
	page.CoverPicture = p.CoverPicture
	page.PostDate = utils.StringToNullTime("2006-01-02 15:04:05", p.PostedTime)
	if page.Slug != p.Slug {
		page.Slug = p.Slug
		page.GUID = fmt.Sprintf("/%s", p.Slug)
	}
	page.ParentID = p.ParentID

	err = dao.Engine.UpdatePage(page, p.Description, p.PageTemplate)
	if err != nil {
		return errno.New(errno.ErrDatabase, err)
	}

	return nil
}

// DeletePost delete post by soft delete
// Note: meta data was reserved
func DeletePost(postType string, articleID uint64) error {
	if err := dao.Engine.DeletePost(postType, articleID); err != nil {
		return errno.New(errno.ErrDatabase, err)
	}

	return nil
}

// TrashPost put the post into the trash by "delete" button
// The different between DeleteArticle and TrashPost is that TrashPost just set the status to deleted
func TrashPost(postID uint64) error {
	if err := dao.Engine.TrashPost(postID); err != nil {
		return errno.New(errno.ErrDatabase, err)
	}

	return nil
}

// RestorePost restore the post which had been put to the trash
// this restore action will set the post as a draft status
func RestorePost(postID uint64) error {
	if err := dao.Engine.RestorePost(postID); err != nil {
		return errno.New(errno.ErrDatabase, err)
	}

	return nil
}
