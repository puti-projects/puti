package handler

import (
	"net/http"
	"strconv"

	"github.com/puti-projects/puti/internal/frontend/service"
	optionCache "github.com/puti-projects/puti/internal/pkg/option"

	"github.com/gin-gonic/gin"
)

// ShowIndex index action
func ShowIndex(c *gin.Context) {
	// get renderer data include basic data
	renderData := getRenderData(c)

	showOnFront := optionCache.Options.Get("show_on_front")
	if showOnFront == "posts" {
		// get params
		currentPage, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

		// get content
		articles, pagination, err := service.GetArticleList(currentPage, "")
		if err != nil {
			// 500
		}

		renderData["Articles"] = articles
		renderData["Pagination"] = pagination.Page

		pagination.SetPageURL("/article")
		renderData["PageURL"] = pagination.PageURL
	} else if showOnFront == "page" {
	} else {
	}

	c.HTML(http.StatusOK, getTheme(c)+"/index.html", renderData)
}
