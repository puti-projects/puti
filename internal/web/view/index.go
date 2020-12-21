package view

import (
	"net/http"
	"strconv"

	"github.com/puti-projects/puti/internal/pkg/cache"
	"github.com/puti-projects/puti/internal/web/service"

	"github.com/gin-gonic/gin"
)

// ShowIndex index action
func ShowIndex(c *gin.Context) {
	// get renderer data include basic data
	renderData := getRenderData(c)

	showOnFront := cache.Options.Get("show_on_front")
	switch showOnFront {
	case "article":
		// get params
		currentPage, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

		// get content
		articles, pagination, err := service.GetArticleList(currentPage, "")
		if err != nil {
			ShowInternalServerError(c)
			return
		}

		renderData["Articles"] = articles

		renderData["Pagination"] = pagination.Page
		pagination.SetPageURL("/article")
		renderData["PageURL"] = pagination.PageURL
	case "page":
	case "knowledge":
		knowledgeList, err := service.SrvEngine.GetKnowledgeList()
		if err != nil {
			ShowInternalServerError(c)
			return
		}
		renderData["Knowledge"] = knowledgeList
	default:
	}

	renderData["Widgets"] = getWidgets()
	renderData["Title"] = renderData["Setting"].(map[string]interface{})["BlogName"].(string) + " - " + renderData["Setting"].(map[string]interface{})["BlogDescription"].(string)
	c.HTML(http.StatusOK, getTheme(c)+"/index.html", renderData)
}
