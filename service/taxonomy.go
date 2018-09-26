package service

import (
	"strconv"

	"puti/model"
)

// TermInfo terms info
type TermInfo struct {
	ID          uint64 `json:"term_id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Pid         uint64 `json:"parent_term_id"`
}

// TreeNode TaxonomyTree's node struct
type TreeNode struct {
	ID           uint64      `json:"id"`
	Name         string      `json:"name"`
	Slug         string      `json:"slug"`
	Description  string      `json:"description"`
	Count        uint64      `json:"count"`
	TermID       uint64      `json:"termID"`
	ParentTermID uint64      `json:"pid"`
	Children     []*TreeNode `json:"children"`
}

// GetTaxonomyList get taxonomy tree by type and return a tree structure
func GetTaxonomyList(taxonomyType string) (taxonomyTree []*TreeNode, err error) {
	termTaxonomy, err := model.GetAllTermsByType(taxonomyType)
	if err != nil {
		return nil, err
	}

	list := GetTaxonomyTree(termTaxonomy, 0)

	return list, nil
}

// GetTaxonomyTree return a taxonomy tree
func GetTaxonomyTree(termTaxonomy []*model.TermTaxonomyModel, pid uint64) []*TreeNode {
	var tree []*TreeNode

	for _, v := range termTaxonomy {
		if pid == v.ParentTermID {
			treeNode := TreeNode{
				ID:           v.ID,
				Name:         v.Term.Name,
				Slug:         v.Term.Slug,
				Description:  v.Term.Description,
				Count:        v.Term.Count,
				TermID:       v.TermID,
				ParentTermID: v.ParentTermID,
			}
			treeNode.Children = GetTaxonomyTree(termTaxonomy, v.ID)
			tree = append(tree, &treeNode)
		}
	}

	return tree
}

// GetTaxonomyInfo get term info by term_id
func GetTaxonomyInfo(termID string) (*TermInfo, error) {
	ID, _ := strconv.Atoi(termID)

	info, err := model.GetTermsInfo(uint64(ID))
	if err != nil {
		return nil, err
	}

	termInfo := &TermInfo{
		ID:          info.ID,
		Name:        info.Term.Name,
		Slug:        info.Term.Slug,
		Description: info.Term.Description,
		Pid:         info.ParentTermID,
	}

	return termInfo, nil
}
