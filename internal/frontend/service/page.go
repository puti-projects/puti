package service

import (
	"html/template"

	"github.com/puti-projects/puti/internal/common/model"
	"github.com/puti-projects/puti/internal/common/utils"
	optionCache "github.com/puti-projects/puti/internal/pkg/option"
)

// GetPageIDBySlug get page ID by slug
func GetPageIDBySlug(slug string) uint64 {
	// get term id
	var pageID uint64
	getpageID := model.DB.Local.Table("pt_post").
		Select("`id`").
		Where("`slug` = ? AND `post_type` = ? AND `parent_id` = ? AND `status` = ? AND `deleted_time` IS NULL", slug, model.PostTypePage, 0, model.PostStatusPublish).
		Row()
	getpageID.Scan(&pageID)

	return pageID
}

// GetPageDetailByID get page detail info by page id
func GetPageDetailByID(pageID uint64) (*model.ShowPageDetail, error) {
	p := &model.PostModel{}
	err := model.DB.Local.Where("id = ? AND post_type = ? AND parent_id = ? AND status =?", pageID, model.PostTypePage, 0, model.PostStatusPublish).First(&p).Error
	if err != nil {
		return nil, err
	}

	siteURL := optionCache.Options.Get("site_url")

	pageDetail := &model.ShowPageDetail{
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

	// get extra data of article
	pm, err := model.GetPostMetaData(pageID)
	if err != nil {
		return nil, err
	}
	for _, meta := range pm {
		pageDetail.MetaData[meta.MetaKey] = meta.MetaValue
	}

	return pageDetail, nil
}
