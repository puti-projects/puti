package service

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/puti-projects/puti/internal/model"
	"github.com/puti-projects/puti/internal/pkg/db"
	"github.com/puti-projects/puti/internal/utils"

	"gorm.io/gorm"
)

// SubjectTreeNode tree struct of subject list
type SubjectTreeNode struct {
	ID          uint64             `json:"id"`
	ParentID    uint64             `json:"parent_id"`
	Name        string             `json:"name"`
	Slug        string             `json:"slug"`
	Description string             `json:"description"`
	Count       uint64             `json:"count"`
	LastUpdated string             `json:"last_updated"`
	Children    []*SubjectTreeNode `json:"children"`
}

// SubjectDetail subject info
type SubjectDetail struct {
	ID               uint64 `json:"id"`
	ParentID         uint64 `json:"parent_id"`
	Name             string `json:"name"`
	Slug             string `json:"slug"`
	Description      string `json:"description"`
	CoverImage       uint64 `json:"cover_image"`
	CoverImageStatus string `json:"cover_image_status"`
	CoverURL         string `json:"cover_url"`
}

// GetSubjectList get subject list by tree struct
func GetSubjectList() ([]*SubjectTreeNode, error) {
	subjects, err := model.GetAllSubjects()
	if err != nil {
		return nil, err
	}

	list := GetSubjectTree(subjects, 0)

	return list, nil
}

// GetSubjectTree return a subject tree
func GetSubjectTree(subjects []*model.SubjectModel, pid uint64) []*SubjectTreeNode {
	var tree []*SubjectTreeNode

	for _, v := range subjects {
		if pid == v.ParentID {
			subjectTreeNode := SubjectTreeNode{
				ID:          v.ID,
				ParentID:    v.ParentID,
				Name:        v.Name,
				Slug:        v.Slug,
				Description: v.Description,
				Count:       v.Count,
			}
			lastUpdatedTime := utils.GetFormatNullTime(&v.LastUpdated, "2006-01-02 15:04:05")
			if lastUpdatedTime == "" {
				subjectTreeNode.LastUpdated = "暂无更新"
			} else {
				subjectTreeNode.LastUpdated = lastUpdatedTime
			}
			subjectTreeNode.Children = GetSubjectTree(subjects, v.ID)
			tree = append(tree, &subjectTreeNode)
		}
	}

	return tree
}

// GetSubjectInfo get subject detail by id
func GetSubjectInfo(subjectID uint64) (*SubjectDetail, error) {
	s, err := model.GetSubjectByID(subjectID)
	if err != nil {
		return nil, err
	}

	var coverImageStatus string
	var coverImageURL string
	if s.CoverImage != 0 {
		m, err := GetMediaByID(s.CoverImage)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			coverImageStatus = "关联封面图不存在，可能已被删除。"
		}
		coverImageURL = m.GUID
	}

	subjectInfo := &SubjectDetail{
		ID:               s.ID,
		ParentID:         s.ParentID,
		Name:             s.Name,
		Slug:             s.Slug,
		Description:      s.Description,
		CoverImage:       s.CoverImage,
		CoverImageStatus: coverImageStatus,
		CoverURL:         coverImageURL,
	}

	return subjectInfo, nil
}

// UpdateSubject udpate subject info
func UpdateSubject(subject *model.SubjectModel) error {
	// begin transcation
	tx := db.DBEngine.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	oldSubject, err := model.GetSubjectByID(subject.ID)
	if err != nil {
		return err
	}

	// if change parent ID
	if oldSubject.ParentID != subject.ParentID {
		if err := updateParentSubjectCount(tx, oldSubject.ParentID, subject.ParentID, oldSubject.Count); err != nil {
			tx.Rollback()
			return err
		}
	}

	oldSubject.ParentID = subject.ParentID
	oldSubject.Name = strings.TrimSpace(subject.Name)
	oldSubject.Slug = strings.TrimSpace(subject.Slug)
	oldSubject.Description = strings.TrimSpace(subject.Description)
	oldSubject.CoverImage = subject.CoverImage

	if err = tx.Model(&model.SubjectModel{}).Save(oldSubject).Error; err != nil {
		tx.Rollback()
		return err
	}

	// commit
	return tx.Commit().Error
}

// updateParentSubjectCount update old parent's and new parent's count
func updateParentSubjectCount(tx *gorm.DB, oldParentID, newParentID, countNum uint64) error {
	// update old parents count
	if oldParentID != 0 {
		if err := updateSubjectCount(tx, oldParentID, -countNum); err != nil {
			return err
		}
	}

	// update new parents count
	if newParentID != 0 {
		if err := updateSubjectCount(tx, newParentID, countNum); err != nil {
			return err
		}
	}

	return nil
}

// updateSubjectCount update subject's count by subject ID and diff count number
func updateSubjectCount(tx *gorm.DB, subjectID uint64, countDiff uint64) error {
	subject := &model.SubjectModel{}
	t := tx.Where("`id` = ?", subjectID).First(&subject)
	if t.Error != nil {
		return t.Error
	}
	subject.Count = subject.Count + countDiff
	if err := tx.Model(&model.SubjectModel{}).Save(subject).Error; err != nil {
		return err
	}

	if subject.ParentID != 0 {
		return updateSubjectCount(tx, subject.ParentID, countDiff)
	}

	return nil
}

// GetArticleSubejct get article's subject by article id
func GetArticleSubejct(articleID uint64) ([]uint64, error) {
	subject, err := model.GetArticleSubject(articleID)
	if err != nil {
		return nil, err
	}

	articleSubject := make([]uint64, 0)
	for _, item := range subject {
		articleSubject = append(articleSubject, item.SubjectID)
	}

	return articleSubject, nil
}

// UpdateSubjectInfoByArticleChange update subject's info (count, last updated time) when creating or updaing the article
// checkout taxonomy's parent and compare it with the subjectIDGroup is in need
func UpdateSubjectInfoByArticleChange(tx *gorm.DB, subjectIDGroup []uint64, countDiff int64, updateLastUpdated bool) (err error) {
	var parentIDs []uint64
	for _, subjectID := range subjectIDGroup {
		parentIDs, err = getSubjectParentID(tx, subjectID, parentIDs)
		if err != nil {
			return err
		}
	}

	if len(parentIDs) != 0 {
		for _, v := range parentIDs {
			inGroup := false
			for _, vv := range subjectIDGroup {
				if vv == v {
					inGroup = true
				}
			}

			if inGroup == false {
				subjectIDGroup = append(subjectIDGroup, v)
			}
		}
	}

	if len(subjectIDGroup) != 0 {
		subjectModel := &model.SubjectModel{}
		updateColumns := map[string]interface{}{}
		if countDiff != 0 {
			updateColumns["count"] = gorm.Expr("count + ?", countDiff)
		}
		if updateLastUpdated {
			updateColumns["last_updated"] = sql.NullTime{Time: time.Now(), Valid: true}
		}

		// exec
		if len(updateColumns) != 0 {
			err = tx.Table(subjectModel.TableName()).Where("`id` IN (?)", subjectIDGroup).Updates(updateColumns).Error
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// getSubjectParentID get all level parents
func getSubjectParentID(tx *gorm.DB, subjectID uint64, parentIDs []uint64) (parentIDGroup []uint64, err error) {
	subject := &model.SubjectModel{}
	err = tx.Where("`id` = ?", subjectID).First(&subject).Error
	if err != nil {
		return nil, err
	}

	if subject.ParentID != 0 {
		parentIDGroup = append(parentIDGroup, subject.ParentID)
		parentIDGroup, err = getSubjectParentID(tx, subject.ParentID, parentIDGroup)
		if err != nil {
			return nil, err
		}
	}

	return parentIDGroup, nil
}

// IfSubjectHasChild check the subject has children or not
func IfSubjectHasChild(subjectID uint64) bool {
	count := model.GetSubjectChildrenNumber(subjectID)
	if count > 0 {
		return true
	}

	return false
}

// DeleteSubject delete subject directly
func DeleteSubject(subjectID uint64) error {
	// begin transcation
	tx := db.DBEngine.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// get subject info
	subject := &model.SubjectModel{}
	t := tx.Where("`id` = ?", subjectID).First(&subject)
	if t.Error != nil {
		return t.Error
	}

	// delete subject
	dSubject := tx.Where("`id` = ?", subjectID).Delete(model.SubjectModel{})
	if err := dSubject.Error; err != nil {
		tx.Rollback()
		return err
	}

	// delete relationship
	dRelation := tx.Where("`subject_id` = ?", subjectID).Delete(model.SubjectRelationshipsModel{})
	if err := dRelation.Error; err != nil {
		tx.Rollback()
		return err
	}

	// update parent count number(if has parent)
	if subject.ParentID > 0 {
		if err := updateSubjectCount(tx, subject.ParentID, -subject.Count); err != nil {
			tx.Rollback()
			return err
		}
	}

	// commit
	return tx.Commit().Error
}
