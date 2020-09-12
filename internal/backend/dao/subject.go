package dao

import (
	"database/sql"
	"strings"
	"time"

	"github.com/puti-projects/puti/internal/model"
	"github.com/puti-projects/puti/internal/pkg/db"

	"gorm.io/gorm"
)

// CreateSubject create a subject
func (d *Dao) CreateSubject(s *model.Subject) error {
	return s.Create(d.db)
}

// GetSubjectByID get subject by id
func (d *Dao) GetSubjectByID(subjectID uint64) (*model.Subject, error) {
	subject := &model.Subject{
		Model: model.Model{ID: subjectID},
	}
	err := subject.GetByID(d.db)
	if err != nil {
		return nil, err
	}

	return subject, nil
}

// UpdateSubject update subject
func (d *Dao) UpdateSubject(subject *model.Subject) error {
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

	oldSubject := &model.Subject{
		Model: model.Model{ID: subject.ID},
	}
	if err := oldSubject.GetByID(d.db); err != nil {
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

	if err := oldSubject.Save(tx); err != nil {
		tx.Rollback()
		return err
	}

	// commit
	return tx.Commit().Error
}

// DeleteSubject delete subject
func (d *Dao) DeleteSubject(subjectID uint64) error {
	// begin transcation
	tx := d.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}

	// get subject info
	subject := &model.Subject{
		Model: model.Model{ID: subjectID},
	}
	if err := subject.GetByID(tx); err != nil {
		return err
	}

	// delete subject
	if err := subject.Delete(tx); err != nil {
		tx.Rollback()
		return err
	}

	// delete relationship
	sr := model.SubjectRelationships{}
	if err := sr.DeleteByCondition(tx, "`subject_id` = ?", []interface{}{subjectID}); err != nil {
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
	subject := &model.Subject{Model: model.Model{ID: subjectID}}
	if err := subject.GetByID(tx); err != nil {
		return err
	}

	subject.Count = subject.Count + countDiff
	if err := subject.Save(tx); err != nil {
		return err
	}

	if subject.ParentID != 0 {
		return updateSubjectCount(tx, subject.ParentID, countDiff)
	}

	return nil
}

// GetArticleSubjectByArticleID get article related subject data
func (d *Dao) GetArticleSubjectByArticleID(articleID uint64) ([]*model.SubjectRelationships, error) {
	sr := &model.SubjectRelationships{}
	subjectRelationships, err := sr.GetAllByObjectID(d.db, articleID)
	if err != nil {
		return nil, err
	}

	return subjectRelationships, nil
}

// updateSubjectInfoByArticleChange update subject's info (count, last updated time) when creating or updaing the article
// checkout taxonomy's parent and compare it with the subjectIDGroup is in need
func updateSubjectInfoByArticleChange(tx *gorm.DB, subjectIDGroup []uint64, countDiff int64, updateLastUpdated bool) (err error) {
	var parentIDGroup []uint64
	for _, subjectID := range subjectIDGroup {
		parentIDGroup, err = getSubjectParentID(tx, subjectID, parentIDGroup)
		if err != nil {
			return err
		}
	}

	if len(parentIDGroup) != 0 {
		for _, v := range parentIDGroup {
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
		updateColumns := map[string]interface{}{}
		if countDiff != 0 {
			updateColumns["count"] = gorm.Expr("count + ?", countDiff)
		}
		if updateLastUpdated {
			updateColumns["last_updated"] = sql.NullTime{Time: time.Now(), Valid: true}
		}

		// exec
		if len(updateColumns) != 0 {
			err = tx.Model(&model.Subject{}).Where("`id` IN (?)", subjectIDGroup).Updates(updateColumns).Error
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// getSubjectParentID get all level parents
func getSubjectParentID(tx *gorm.DB, subjectID uint64, parentIDs []uint64) (parentIDGroup []uint64, err error) {
	subject := &model.Subject{
		Model: model.Model{ID: subjectID},
	}
	err = subject.GetByID(tx)
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

// CheckSubjectNameExist check the subject name
func (d *Dao) CheckSubjectNameExist(subjectID uint64, name string) bool {
	subject := &model.Subject{
		Model: model.Model{ID: subjectID},
		Name:  name,
	}
	return subject.CheckSubjectNameExist(d.db)
}

// IfSubjectHasChild calcuelate the total number of subject's children
func (d *Dao) IfSubjectHasChild(subjectID uint64) bool {
	s := &model.Subject{}
	return s.IfSubjectHasChild(d.db, subjectID)
}

// GetAllSubjects get all subjects
func (d *Dao) GetAllSubjects() ([]*model.Subject, error) {
	s := &model.Subject{}
	return s.GetAll(d.db)
}
