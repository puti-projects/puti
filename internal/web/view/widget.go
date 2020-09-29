package view

import (
	"github.com/puti-projects/puti/internal/pkg/logger"
	"github.com/puti-projects/puti/internal/web/service"
)

func widgetLatestArticles(showNums int) []*service.ShowWidgetLatestArticles {
	list, err := service.GetLatestArticlesList(showNums)
	if err != nil {
		logger.Errorf("get latest article list failed. %s", err)
		return nil
	}

	return list
}

func widgetCategoryList() []*service.ShowWidgetCategoryTreeNode {
	list, err := service.GetcategoryList()
	if err != nil {
		logger.Errorf("get category list failed. %s", err)
		return nil
	}

	return list
}
