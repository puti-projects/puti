package service

import (
	"html/template"

	"github.com/puti-projects/puti/internal/model"
	"github.com/puti-projects/puti/internal/pkg/db"
	optionCache "github.com/puti-projects/puti/internal/pkg/option"
	"github.com/puti-projects/puti/internal/utils"
)

// GetPageIDBySlug get page ID by slug
func GetPageIDBySlug(slug string) uint64 {
	// get term id
	var pageID uint64
	getpageID := db.Engine.Table("pt_post").
		Select("`id`").
		Where("`slug` = ? AND `post_type` = ? AND `parent_id` = ? AND `status` = ? AND `deleted_time` IS NULL", slug, model.PostTypePage, 0, model.PostStatusPublish).
		Row()
	_ = getpageID.Scan(&pageID)

	return pageID
}

// GetPageDetailByID get page detail info by page id
func GetPageDetailByID(pageID uint64) (*ShowPageDetail, error) {
	p := &model.Post{}
	err := db.Engine.Where("id = ? AND post_type = ? AND parent_id = ? AND status =?", pageID, model.PostTypePage, 0, model.PostStatusPublish).First(&p).Error
	if err != nil {
		return nil, err
	}

	siteURL := optionCache.Options.Get("site_url")

	pageDetail := &ShowPageDetail{
		ID:            p.ID,
		Title:         p.Title,
		ContentHTML:   template.HTML(p.ContentHTML),
		CommentStatus: p.CommentStatus,
		GUID:          siteURL + p.GUID,
		CommentCount:  p.CommentCount,
		ViewCount:     p.ViewCount,
		PostedTime:    utils.GetFormatNullTime(&p.PostDate, "2006-01-02 15:04"),
		MetaData:      make(map[string]interface{}),
	}

	// get extra data of page
	meta := &model.PostMeta{PostID: pageID}
	pm, err := meta.GetAllByPostID(db.Engine)
	if err != nil {
		return nil, err
	}
	for _, meta := range pm {
		pageDetail.MetaData[meta.MetaKey] = meta.MetaValue
	}

	return pageDetail, nil
}
