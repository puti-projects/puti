package service

import (
	"github.com/puti-projects/puti/internal/common/model"
	"github.com/puti-projects/puti/internal/common/utils"
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
