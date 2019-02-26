package service

import (
	"sort"

	"github.com/puti-projects/puti/internal/common/config"
	"github.com/puti-projects/puti/internal/common/model"
	"github.com/puti-projects/puti/internal/common/utils"
	"github.com/puti-projects/puti/internal/pkg/logger"
)

// GetArchive get archive list and list sort
func GetArchive() (map[string]map[string][]*model.ShowArchive, []string, map[string][]string, error) {
	archives := []model.PostModel{}

	where := "`post_type` = ? AND `parent_id` = ? AND `status` = ?"
	whereArgs := []interface{}{model.PostTypeArticle, 0, model.PostStatusPublish}
	postModel := &model.PostModel{}
	rows, err := model.DB.Local.Table(postModel.TableName()).
		Select("`id`, `title`, `guid`, `comment_count`, `view_count`, `posted_time`").
		Where(where, whereArgs...).
		Order("`posted_time` DESC").
		Rows()
	defer rows.Close()
	for rows.Next() {
		var archive model.PostModel
		// ScanRows scan a row into archive
		model.DB.Local.ScanRows(rows, &archive)
		archives = append(archives, archive)
	}
	if err != nil {
		logger.Errorf("get all articles failed. %s", err)
		return nil, nil, nil, err
	}

	dataMap := map[string]map[string][]*model.ShowArchive{}
	sortYear := []string{}
	sortMonth := map[string][]string{}
	for _, v := range archives {
		postedYear := v.PostDate.In(config.TimeLoc()).Format("2006")
		postedMonth := v.PostDate.In(config.TimeLoc()).Format("01")

		_, existYear := dataMap[postedYear]
		if !existYear {
			dataMap[postedYear] = make(map[string][]*model.ShowArchive)
			sortYear = append(sortYear, postedYear)
		}

		_, existMonth := dataMap[postedYear][postedMonth]
		if !existMonth {
			dataMap[postedYear][postedMonth] = make([]*model.ShowArchive, 0)
			if _, existSortYear := sortMonth[postedYear]; !existSortYear {
				sortMonth[postedYear] = []string{}
			}
			sortMonth[postedYear] = append(sortMonth[postedYear], postedMonth)
		}

		article := &model.ShowArchive{
			ID:           v.ID,
			Title:        v.Title,
			GUID:         v.GUID,
			CommentCount: v.CommentCount,
			ViewCount:    v.ViewCount,
			PostedTime:   utils.GetFormatTime(&v.PostDate, "2006-01-02 15:04"),
			PostedDay:    v.PostDate.In(config.TimeLoc()).Format("02"),
		}

		dataMap[postedYear][postedMonth] = append(dataMap[postedYear][postedMonth], article)
	}

	// sort those sort maps
	sort.Sort(sort.Reverse(sort.StringSlice(sortYear)))
	for _, monthSort := range sortMonth {
		sort.Sort(sort.Reverse(sort.StringSlice(monthSort)))
	}

	return dataMap, sortYear, sortMonth, nil
}
