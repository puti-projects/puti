package service

import "github.com/puti-projects/puti/internal/common/model"

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
