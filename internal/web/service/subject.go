package service

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/puti-projects/puti/internal/model"
	"github.com/puti-projects/puti/internal/pkg/db"
	"github.com/puti-projects/puti/internal/utils"
)

// SubjectList subject list
type SubjectList struct {
	Lock  *sync.Mutex
	IDMap map[uint64]*ShowSubjectList
}

// SubjectInfoResult use for scan subject info
type SubjectInfoResult struct {
	ID            uint64
	ParentID      uint64
	Name          string
	Slug          string
	Description   string
	CoverImageURL string
	Count         string
}

// ChildrenSubjectsResult use for scan
type ChildrenSubjectsResult struct {
	ID            uint64
	ParentID      uint64
	Name          string
	Slug          string
	Description   string
	CoverImageURL string
	Count         uint64
	LastUpdated   sql.NullTime
}

// GetChildrenSubjects get subject's all children
// If the parentID is 0, get all top subjects
func GetChildrenSubjects(parentID uint64) (subjectResult []*ShowSubjectList, err error) {
	subjects := make([]*ChildrenSubjectsResult, 0)
	subjectModel := &model.Subject{}
	rows, err := db.Engine.Table(subjectModel.TableName()+" s").
		Select("s.`id`, s.`parent_id`, s.`name`, s.`slug`, s.`description`, r.`guid` as cover_image_url, s.`count`, s.`last_updated`").
		Joins("LEFT JOIN pt_resource r ON r.id = s.`cover_image`").
		Where("s.`parent_id` = ? AND s.`deleted_time` is null", parentID).
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		subject := &ChildrenSubjectsResult{}
		// ScanRows scan a row into subject
		if err := db.Engine.ScanRows(rows, &subject); err != nil {
			return nil, err
		}

		subjects = append(subjects, subject)
	}

	subjectResult = make([]*ShowSubjectList, 0)
	var ids []uint64
	for _, subject := range subjects {
		ids = append(ids, subject.ID)
	}

	wg := sync.WaitGroup{}
	subjectList := SubjectList{
		Lock:  new(sync.Mutex),
		IDMap: make(map[uint64]*ShowSubjectList, len(subjects)),
	}

	finished := make(chan bool, 1)

	for _, s := range subjects {
		wg.Add(1)
		go func(s *ChildrenSubjectsResult) {
			defer wg.Done()

			subjectList.Lock.Lock()
			defer subjectList.Lock.Unlock()
			subjectList.IDMap[s.ID] = &ShowSubjectList{
				ID:                s.ID,
				URL:               "/subject/" + s.Slug,
				Name:              s.Name,
				Slug:              s.Slug,
				Description:       s.Description,
				CoverImageURL:     s.CoverImageURL,
				Count:             s.Count,
				LastUpdated:       utils.GetFormatNullTime(&s.LastUpdated, "2006-01-02 15:04"),
				SubLastUpdatedDay: getDifferDayBetweenLastUpdatedTimeAndNow(&s.LastUpdated),
			}
		}(s)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	<-finished

	for _, id := range ids {
		subjectResult = append(subjectResult, subjectList.IDMap[id])
	}

	return
}

// getDifferDayBetweenLastUpdatedTimeAndNow calculate updated how many days ago
func getDifferDayBetweenLastUpdatedTimeAndNow(lastUpdatedTime *sql.NullTime) string {
	if lastUpdatedTime.Valid {
		day := utils.SubNullTimeUnitlNowAsDay(lastUpdatedTime.Time)
		if day < 1 {
			return "24 小时内有更新"
		}

		str := fmt.Sprintf("%v 天前更新", uint64(day))
		return str
	}

	return "暂无更新"
}

// GetSubjectInfoBySlug get subject info by the unique slug
func GetSubjectInfoBySlug(subjectSlug string) (*ShowSubjectInfo, error) {
	subjectResult := &SubjectInfoResult{}
	subjectModel := &model.Subject{}
	result := db.Engine.Table(subjectModel.TableName()+" s").
		Select("s.`id`, s.`parent_id`, s.`name`, s.`slug`, s.`description`, r.`guid` as cover_image_url, s.`count`").
		Joins("LEFT JOIN pt_resource r ON r.id = s.`cover_image`").
		Where("s.`slug` = ? AND s.`deleted_time` is null", subjectSlug).
		Find(&subjectResult)
	if err := result.Error; err != nil {
		return nil, err
	}

	// siteURL := optionCache.Options.Get("site_url")
	subjectInfo := &ShowSubjectInfo{
		ID:            subjectResult.ID,
		ParentURL:     "/subject",
		Name:          subjectResult.Name,
		Slug:          subjectResult.Slug,
		Description:   subjectResult.Description,
		CoverImageURL: subjectResult.CoverImageURL,
		Count:         subjectResult.Count,
	}

	// get parent url
	if subjectResult.ParentID > 0 {
		parent := model.Subject{Model: model.Model{ID: subjectResult.ParentID}}
		if err := parent.GetByID(db.Engine); err != nil {
			return nil, err
		}
		subjectInfo.ParentURL = "/subject/" + parent.Slug
	}

	return subjectInfo, nil
}
