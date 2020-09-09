package handler

import (
	"github.com/puti-projects/puti/internal/routers/middleware/view"

	"github.com/gin-gonic/gin"
)

// getDataModel consume renderer data
func getRenderData(c *gin.Context) view.RenderData {
	renderData, _ := c.Get("renderData")

	return *(renderData.(*view.RenderData))
}

// getTheme return current theme name
func getTheme(c *gin.Context) string {
	renderData := getRenderData(c)

	return renderData["Setting"].(map[string]interface{})["CurrentTheme"].(string)
}

// getSiteURL return current site url from setting
func getSiteURL(c *gin.Context) string {
	renderData := getRenderData(c)

	return renderData["Setting"].(map[string]interface{})["SiteUrl"].(string)
}

// getWidgets load setting widgets into renderData
func getWidgets() map[string]interface{} {
	widgetMap := map[string]interface{}{}

	widgetMap["LatestArticles"] = widgetLatestArticles(6)
	widgetMap["CategoryList"] = widgetCategoryList()

	return widgetMap
}
