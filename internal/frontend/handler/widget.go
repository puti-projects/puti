package handler

import (
	"github.com/puti-projects/puti/internal/frontend/service"
	"github.com/puti-projects/puti/internal/model"
	"github.com/puti-projects/puti/internal/pkg/logger"
)

func widgetLatestArticles(showNums int) []*model.ShowWidgetLatestArticles {
	list, err := service.GetLatestArticlesList(showNums)
	if err != nil {
		logger.Errorf("get latest article list failed. %s", err)
		return nil
	}

	return list
}

func widgetCategoryList() []*model.ShowWidgetCategoryTreeNode {
	list, err := service.GetcategoryList()
	if err != nil {
		logger.Errorf("get category list failed. %s", err)
		return nil
	}

	return list
}
