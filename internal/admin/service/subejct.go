package service

import (
	"errors"

	"github.com/puti-projects/puti/internal/model"
	"github.com/puti-projects/puti/internal/pkg/errno"
	"github.com/puti-projects/puti/internal/utils"

	"gorm.io/gorm"
)

// SubjectCreateRequest struct bind to create subject
type SubjectCreateRequest struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	ParentID    uint64 `json:"parent_id"`
	CoverImage  uint64 `json:"cover_image"`
	Description string `json:"description"`
}

// SubjectUpdateRequest struct bind to update subject
type SubjectUpdateRequest struct {
	ID          uint64 `json:"ID"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	ParentID    uint64 `json:"parent_id"`
	CoverImage  uint64 `json:"cover_image"`
	Description string `json:"description"`
}

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

// CreateSubject create a subject
func (svc Service) CreateSubject(r *SubjectCreateRequest) error {
	s := &model.Subject{
		ParentID:    r.ParentID,
		Name:        r.Name,
		Slug:        r.Slug,
		Description: r.Description,
		CoverImage:  r.CoverImage,
	}

	if err := svc.dao.CreateSubject(s); err != nil {
		return errno.New(errno.ErrDatabase, err)
	}

	return nil
}

// CheckSubjectNameExist check the subject name
func (svc Service) CheckSubjectNameExist(subjectID uint64, name string) bool {
	return svc.dao.CheckSubjectNameExist(subjectID, name)
}

// GetSubjectList get subject list by tree struct
func (svc Service) GetSubjectList() ([]*SubjectTreeNode, error) {
	subjects, err := svc.dao.GetAllSubjects()
	if err != nil {
		return nil, errno.New(errno.ErrDatabase, err)
	}

	list := getSubjectTree(subjects, 0)

	return list, nil
}

// GetSubjectTree return a subject tree
func getSubjectTree(subjects []*model.Subject, pid uint64) []*SubjectTreeNode {
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
			subjectTreeNode.Children = getSubjectTree(subjects, v.ID)
			tree = append(tree, &subjectTreeNode)
		}
	}

	return tree
}

// GetSubjectInfo get subject detail by id
func (svc Service) GetSubjectInfo(subjectID uint64) (*SubjectDetail, error) {
	s, err := svc.dao.GetSubjectByID(subjectID)
	if err != nil {
		return nil, err
	}

	var coverImageStatus string
	var coverImageURL string
	if s.CoverImage != 0 {
		m, err := svc.GetMediaByID(s.CoverImage)
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

// UpdateSubject update subject info
func (svc Service) UpdateSubject(r *SubjectUpdateRequest) error {
	subject := &model.Subject{
		Model: model.Model{ID: r.ID},

		ParentID:    r.ParentID,
		Name:        r.Name,
		Slug:        r.Slug,
		Description: r.Description,
		CoverImage:  r.CoverImage,
	}

	if err := svc.dao.UpdateSubject(subject); err != nil {
		return errno.New(errno.ErrDatabase, err)
	}

	return nil
}

// GetArticleSubjectID get article's subject by article id
func (svc Service) GetArticleSubjectID(articleID uint64) ([]uint64, error) {
	subjectRelation, err := svc.dao.GetArticleSubjectByArticleID(articleID)
	if err != nil {
		return nil, err
	}

	articleSubject := make([]uint64, 0)
	for _, item := range subjectRelation {
		articleSubject = append(articleSubject, item.SubjectID)
	}

	return articleSubject, nil
}

// checkIfSubjectCanDelete check the subject has children or not
func (svc Service) checkIfSubjectCanDelete(subjectID uint64) error {
	if ifHasChild := svc.dao.IfSubjectHasChild(subjectID); ifHasChild == true {
		return errno.New(errno.ErrValidation, nil).Add("subject has children and can not be deleted")
	}

	return nil
}

// DeleteSubject delete subject directly
func (svc Service) DeleteSubject(subjectID uint64) error {
	// check
	if err := svc.checkIfSubjectCanDelete(subjectID); err != nil {
		return err
	}

	if err := svc.dao.DeleteSubject(subjectID); err != nil {
		return errno.New(errno.ErrDatabase, err)
	}

	return nil
}
