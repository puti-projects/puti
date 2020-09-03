package service

import (
	"github.com/puti-projects/puti/internal/model"
	"github.com/puti-projects/puti/internal/pkg/db"
)

// GetcategoryList get category tree for widget
func GetcategoryList() (taxonomyTree []*model.ShowWidgetCategoryTreeNode, err error) {
	var termTaxonomy []*model.TermTaxonomyModel
	err = db.DBEngine.Where("taxonomy = ? AND term_id != ?", "category", model.DefaultUnCategorizedID).Preload("Term").Find(&termTaxonomy).Error
	if err != nil {
		return nil, err
	}

	list := getTaxonomyTree(termTaxonomy, 0)

	return list, nil
}

// getTaxonomyTree return a taxonomy tree
func getTaxonomyTree(termTaxonomy []*model.TermTaxonomyModel, pid uint64) []*model.ShowWidgetCategoryTreeNode {
	var tree []*model.ShowWidgetCategoryTreeNode

	for _, v := range termTaxonomy {
		if pid == v.ParentTermID {
			treeNode := model.ShowWidgetCategoryTreeNode{
				TermID: v.TermID,
				Name:   v.Term.Name,
				Slug:   v.Term.Slug,
				Count:  v.Term.Count,
				URL:    "/category/" + v.Term.Slug, // TODO could be setting as a param
			}
			treeNode.Children = getTaxonomyTree(termTaxonomy, v.TermID)
			tree = append(tree, &treeNode)
		}
	}

	return tree
}
