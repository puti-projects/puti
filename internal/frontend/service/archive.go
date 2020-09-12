package service

import (
	"sort"

	"github.com/puti-projects/puti/internal/model"
	"github.com/puti-projects/puti/internal/pkg/db"
	"github.com/puti-projects/puti/internal/pkg/logger"
	"github.com/puti-projects/puti/internal/utils"
)

// GetArchive get archive list and list sort
func GetArchive() (map[string]map[string][]*ShowArchive, []string, map[string][]string, error) {
	archives := []model.Post{}

	where := "`post_type` = ? AND `parent_id` = ? AND `status` = ?"
	whereArgs := []interface{}{model.PostTypeArticle, 0, model.PostStatusPublish}
	postModel := &model.Post{}
	rows, err := db.DBEngine.Table(postModel.TableName()).
		Select("`id`, `title`, `guid`, `comment_count`, `view_count`, `posted_time`").
		Where(where, whereArgs...).
		Order("`posted_time` DESC").
		Rows()
	if err != nil {
		logger.Errorf("get all articles failed. %s", err)
		return nil, nil, nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var archive model.Post
		// ScanRows scan a row into archive
		db.DBEngine.ScanRows(rows, &archive)
		archives = append(archives, archive)
	}

	dataMap := map[string]map[string][]*ShowArchive{}
	sortYear := []string{}
	sortMonth := map[string][]string{}
	for _, v := range archives {
		postedYear := utils.GetFormatNullTime(v.PostDate, "2006")
		postedMonth := utils.GetFormatNullTime(v.PostDate, "01")

		_, existYear := dataMap[postedYear]
		if !existYear {
			dataMap[postedYear] = make(map[string][]*ShowArchive)
			sortYear = append(sortYear, postedYear)
		}

		_, existMonth := dataMap[postedYear][postedMonth]
		if !existMonth {
			dataMap[postedYear][postedMonth] = make([]*ShowArchive, 0)
			if _, existSortYear := sortMonth[postedYear]; !existSortYear {
				sortMonth[postedYear] = []string{}
			}
			sortMonth[postedYear] = append(sortMonth[postedYear], postedMonth)
		}

		article := &ShowArchive{
			ID:           v.ID,
			Title:        v.Title,
			GUID:         v.GUID,
			CommentCount: v.CommentCount,
			ViewCount:    v.ViewCount,
			PostedTime:   utils.GetFormatNullTime(v.PostDate, "2006-01-02 15:04"),
			PostedDay:    utils.GetFormatNullTime(v.PostDate, "02"),
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
