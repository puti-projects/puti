package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ShowNotFound 404 handler
func ShowNotFound(c *gin.Context) {
	// get renderer data include basic data
	renderData := getRenderData(c)

	renderData["code"] = "404"
	renderData["message"] = "Sorry! We can't seem to find the page you're looking for."

	renderData["Widgets"] = getWidgets()
	renderData["Title"] = "404 - " + renderData["Setting"].(map[string]interface{})["BlogName"].(string)
	c.HTML(http.StatusNotFound, getTheme(c)+"/error.html", renderData)
}
