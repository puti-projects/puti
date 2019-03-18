package service

import "github.com/puti-projects/puti/internal/common/model"

// DashdoardData some statistics index
type DashdoardData struct {
	TotalViews    uint64
	TotalComments uint64
	TotalArticles uint64
	TotalMedia    uint64
}

// GetDashdoardStatisticsData get dashdoard statistics data
func GetDashdoardStatisticsData() (*DashdoardData, error) {
	var totalViews, totalComments, totalArticles, totalMedia uint64
	postModel := &model.PostModel{}
	totalViewsRow := model.DB.Local.Table(postModel.TableName()).Where("`status` != ? AND `deleted_time` is null", "deleted").
		Select("sum(`view_count`) as total_views").
		Row()
	totalViewsRow.Scan(&totalViews)

	if err := model.DB.Local.Table(postModel.TableName()).
		Where("`post_type` = ? AND `status` != ? AND `deleted_time` is null", model.PostTypeArticle, "deleted").
		Count(&totalArticles).
		Error; err != nil {
		return nil, err
	}

	mediaModel := &model.MediaModel{}
	if err := model.DB.Local.Table(mediaModel.TableName()).Where("`deleted_time` is null").Count(&totalMedia).Error; err != nil {
		return nil, err
	}

	var dashdoardData *DashdoardData
	dashdoardData = &DashdoardData{
		TotalViews:    totalViews,
		TotalComments: totalComments,
		TotalArticles: totalArticles,
		TotalMedia:    totalMedia,
	}

	return dashdoardData, nil
}
