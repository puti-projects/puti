package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/puti-projects/puti/internal/frontend/service"

	"github.com/gin-gonic/gin"
)

// ShowArticleList article list handle
func ShowArticleList(c *gin.Context) {
	// get renderer data include basic data
	renderData := getRenderData(c)

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

	renderData["Widgets"] = getWidgets()
	renderData["Title"] = "文章" + " - " + renderData["Setting"].(map[string]interface{})["BlogName"].(string)
	c.HTML(http.StatusOK, getTheme(c)+"/articles.html", renderData)
}

// ShowCategoryArticleList handle article list by category
func ShowCategoryArticleList(c *gin.Context) {
	// get renderer data include basic data
	renderData := getRenderData(c)

	// get params
	taxonomySlug := c.Param("slug")
	currentPage, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	// get content
	termName, articles, pagination, err := service.GetArticleListByTaxonomy(currentPage, "category", taxonomySlug, "")
	if err != nil {
		ShowInternalServerError(c)
		return
	}

	renderData["Articles"] = articles

	renderData["Pagination"] = pagination.Page
	pagination.SetPageURL("/category/" + taxonomySlug)
	renderData["PageURL"] = pagination.PageURL

	renderData["Widgets"] = getWidgets()
	renderData["Title"] = termName + " - 分类 - " + renderData["Setting"].(map[string]interface{})["BlogName"].(string)
	c.HTML(http.StatusOK, getTheme(c)+"/articles.html", renderData)
}

// ShowTagArticleList handle article list by tag
func ShowTagArticleList(c *gin.Context) {
	// get renderer data include basic data
	renderData := getRenderData(c)

	// get params
	taxonomySlug := c.Param("slug")
	currentPage, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	// get content
	termName, articles, pagination, err := service.GetArticleListByTaxonomy(currentPage, "tag", taxonomySlug, "")
	if err != nil {
		ShowInternalServerError(c)
		return
	}

	renderData["Articles"] = articles

	renderData["Pagination"] = pagination.Page
	pagination.SetPageURL("/tag/" + taxonomySlug)
	renderData["PageURL"] = pagination.PageURL

	renderData["Widgets"] = getWidgets()
	renderData["Title"] = termName + " - 标签 - " + renderData["Setting"].(map[string]interface{})["BlogName"].(string)
	c.HTML(http.StatusOK, getTheme(c)+"/articles.html", renderData)
}

// ShowArticleDetail handle article datail
func ShowArticleDetail(c *gin.Context) {
	renderData := getRenderData(c)

	// get params
	articleID := strings.Split(c.Param("id"), ".")[0]

	articleDetail, err := service.GetArticleDetailByID(articleID)
	if err != nil {
		ShowInternalServerError(c)
		return
	}

	renderData["Article"] = articleDetail

	renderData["Title"] = articleDetail.Title + " - " + renderData["Setting"].(map[string]interface{})["BlogName"].(string)
	c.HTML(http.StatusOK, getTheme(c)+"/article-detail.html", renderData)
}
