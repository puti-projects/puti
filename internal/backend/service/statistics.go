package service

import (
	"github.com/puti-projects/puti/internal/model"
	"github.com/puti-projects/puti/internal/pkg/db"
)

// DashboardData some statistics index
type DashboardData struct {
	TotalViews    uint64
	TotalComments uint64
	TotalArticles uint64
	TotalMedia    uint64
}

// GetDashboardStatisticsData get dashboard statistics data
// TODO store in cache first
func GetDashboardStatisticsData() (*DashboardData, error) {
	var totalViews, totalComments, totalArticles, totalMedia uint64
	postModel := &model.PostModel{}
	totalViewsRow := db.DBEngine.Table(postModel.TableName()).Where("`status` != ? AND `deleted_time` is null", "deleted").
		Select("sum(`view_count`) as total_views").
		Row()
	_ = totalViewsRow.Scan(&totalViews)

	if err := db.DBEngine.Table(postModel.TableName()).
		Where("`post_type` = ? AND `status` != ? AND `deleted_time` is null", model.PostTypeArticle, "deleted").
		Count(&totalArticles).
		Error; err != nil {
		return nil, err
	}

	mediaModel := &model.MediaModel{}
	if err := db.DBEngine.Table(mediaModel.TableName()).Where("`deleted_time` is null").Count(&totalMedia).Error; err != nil {
		return nil, err
	}

	var dashboardData *DashboardData
	dashboardData = &DashboardData{
		TotalViews:    totalViews,
		TotalComments: totalComments,
		TotalArticles: totalArticles,
		TotalMedia:    totalMedia,
	}

	return dashboardData, nil
}
