package handler

import (
	"github.com/puti-projects/puti/internal/frontend/middleware"

	"github.com/gin-gonic/gin"
)

// getDataModel consume renderer data
func getRenderData(c *gin.Context) middleware.RenderData {
	renderData, _ := c.Get("renderData")
	return *(renderData.(*middleware.RenderData))
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
