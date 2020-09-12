package web

import (
	"github.com/puti-projects/puti/internal/routers/middleware/web"

	"github.com/gin-gonic/gin"
)

// getDataModel consume renderer data
func getRenderData(c *gin.Context) web.RenderData {
	renderData, _ := c.Get("renderData")

	return *(renderData.(*web.RenderData))
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
