package service

import (
	"github.com/puti-projects/puti/internal/admin/dao"
	"github.com/puti-projects/puti/internal/model"
	"github.com/puti-projects/puti/internal/pkg/errno"
)

// TaxonomyCreateRequest struct to crate taxonomy include category and tag
type TaxonomyCreateRequest struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	ParentID    uint64 `json:"parentId"`
	Taxonomy    string `json:"taxonomy"` // category or tag
}

// TaxonomyUpdateRequest param struct to update taxonomy include category and tag
type TaxonomyUpdateRequest struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	ParentID    uint64 `json:"parentId"`
	Taxonomy    string `json:"taxonomy"` // category or tag
}

// TermInfo terms info
type TermInfo struct {
	ID          uint64 `json:"term_id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Pid         uint64 `json:"parent_term_id"`
	Level       uint64 `json:"level"`
}

// TreeNode TaxonomyTree's node struct
type TreeNode struct {
	ID           uint64      `json:"id"`
	Name         string      `json:"name"`
	Slug         string      `json:"slug"`
	Description  string      `json:"description"`
	Count        uint64      `json:"count"`
	TermID       uint64      `json:"term_id"`
	ParentTermID uint64      `json:"pid"`
	Level        uint64      `json:"level"`
	Children     []*TreeNode `json:"children"`
}

//CreateTaxonomy create term taxonomy
func CreateTaxonomy(r *TaxonomyCreateRequest) error {
	if r.Slug == "" {
		r.Slug = r.Name
	}

	level, err := dao.Engine.GetTaxonomyLevel(r.ParentID, r.Taxonomy)
	if err != nil {
		return err
	}

	termTaxonomy := &model.TermTaxonomy{
		Term: model.Term{
			Name:        r.Name,
			Slug:        r.Slug,
			Description: r.Description,
			Count:       0,
		},
		ParentTermID: r.ParentID,
		Level:        level + 1,
		Taxonomy:     r.Taxonomy,
		TermGroup:    0,
	}

	if err := dao.Engine.CreateTaxonomy(termTaxonomy); err != nil {
		return errno.New(errno.ErrDatabase, err)
	}

	return nil
}

// GetTaxonomyList get taxonomy tree by type and return a tree structure
func GetTaxonomyList(taxonomyType string) (taxonomyTree []*TreeNode, err error) {
	allTermTaxonomy, err := dao.Engine.GetAllByType(taxonomyType)
	if err != nil {
		return nil, err
	}

	// pid is 0 means level 1; begin from level 1
	list := GetTaxonomyTree(allTermTaxonomy, 0)

	return list, nil
}

// GetTaxonomyTree return a taxonomy tree
func GetTaxonomyTree(allTermTaxonomy []*model.TermTaxonomy, pid uint64) []*TreeNode {
	var tree []*TreeNode

	for _, v := range allTermTaxonomy {
		// get all terms in this level as treeNode
		if pid == v.ParentTermID {
			treeNode := TreeNode{
				ID:           v.ID,
				Name:         v.Term.Name,
				Slug:         v.Term.Slug,
				Description:  v.Term.Description,
				Count:        v.Term.Count,
				TermID:       v.TermID,
				ParentTermID: v.ParentTermID,
				Level:        v.Level,
			}
			// get their children
			treeNode.Children = GetTaxonomyTree(allTermTaxonomy, v.TermID)
			tree = append(tree, &treeNode)
		}
	}

	return tree
}

// GetTaxonomyInfo get term info by term_id
func GetTaxonomyInfo(termID uint64) (*TermInfo, error) {
	termTaxonomy, err := dao.Engine.GetTermTaxonomyByTermID(termID)
	if err != nil {
		return nil, err
	}

	termInfo := &TermInfo{
		ID:          termTaxonomy.ID,
		Name:        termTaxonomy.Term.Name,
		Slug:        termTaxonomy.Term.Slug,
		Description: termTaxonomy.Term.Description,
		Pid:         termTaxonomy.ParentTermID,
		Level:       termTaxonomy.Level,
	}
	return termInfo, nil
}

// UpdateTaxonomy update term and term taxonomy
func UpdateTaxonomy(r *TaxonomyUpdateRequest, termID uint64) error {
	termTaxonomy := &model.TermTaxonomy{
		Term: model.Term{
			ID:          termID,
			Name:        r.Name,
			Slug:        r.Slug,
			Description: r.Description,
		},
		TermID:       termID,
		ParentTermID: r.ParentID,
	}

	// Update changed fields.
	if err := dao.Engine.UpdateTaxonomy(termTaxonomy, r.Taxonomy); err != nil {
		return errno.New(errno.ErrDatabase, err)
	}

	return nil
}

// IfTaxonomyHasChild check the taxonomy has children or not
func IfTaxonomyHasChild(termID uint64, taxonomyType string) (bool, error) {
	count, err := dao.Engine.GetTermChildrenNumber(termID, taxonomyType)
	if err != nil {
		return true, errno.New(errno.ErrDatabase, err)
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

// CheckTaxonomyNameExist check if taxonomy name are aleady exist
func CheckTaxonomyNameExist(name, taxonomy string) bool {
	return dao.Engine.CheckTaxonomyNameExist(name, taxonomy)
}

// DeleteTaxonomy delete term directly
func DeleteTaxonomy(termID uint64, taxonomyType string) error {
	return dao.Engine.DeleteTaxonomy(termID, taxonomyType)
}
