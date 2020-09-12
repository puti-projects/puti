package web

import (
	"net/http"
	"strings"

	"github.com/puti-projects/puti/internal/frontend/service"

	"github.com/gin-gonic/gin"
)

// ShowNotFound 404 handler
func ShowNotFound(c *gin.Context) {
	if !strings.HasPrefix(c.Request.RequestURI, "/themes") {
		// no static, and get slug
		slug := strings.TrimLeft(c.Request.RequestURI, "/")
		if !strings.Contains(slug, "/") {
			if pageID := service.GetPageIDBySlug(slug); pageID > 0 {
				ShowPageDetail(c, pageID)
				return
			}
		}
	}

	// get renderer data include basic data
	renderData := getRenderData(c)

	renderData["code"] = "404"
	renderData["message"] = "Sorry! We can't seem to find the page you're looking for."

	renderData["Widgets"] = getWidgets()
	renderData["Title"] = "404 - " + renderData["Setting"].(map[string]interface{})["BlogName"].(string)
	c.HTML(http.StatusNotFound, getTheme(c)+"/error.html", renderData)
}

// ShowInternalServerError 500 handler
func ShowInternalServerError(c *gin.Context) {
	// get renderer data include basic data
	renderData := getRenderData(c)

	renderData["code"] = "500"
	renderData["message"] = "Something is not quite right. We will be back soon!"

	renderData["Widgets"] = getWidgets()
	renderData["Title"] = "500 - " + renderData["Setting"].(map[string]interface{})["BlogName"].(string)
	c.HTML(http.StatusInternalServerError, getTheme(c)+"/error.html", renderData)
}
