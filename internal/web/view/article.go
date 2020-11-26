package view

import (
	"errors"
	"github.com/puti-projects/puti/internal/pkg/config"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/puti-projects/puti/internal/pkg/counter"
	"github.com/puti-projects/puti/internal/web/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	ID, _ := strconv.Atoi(articleID)
	aID := uint64(ID)

	// check cache
	if data, exist := service.SrvEngine.GetCache(config.CacheArticleDetailPrefix + articleID); exist {
		s := &map[string]interface{}{}
		service.SrvEngine.JSONUnmarshal(data, &s)

		articleDetail := (*s)["Article"].(map[string]interface{})
		articleDetail["ContentHTML"] = template.HTML(articleDetail["ContentHTML"].(string))

		renderData["Article"] = articleDetail
		renderData["LastArticle"] = (*s)["LastArticle"]
		renderData["NextArticle"] = (*s)["NextArticle"]
		renderData["Title"] = articleDetail["Title"].(string) + " - " + renderData["Setting"].(map[string]interface{})["BlogName"].(string)
	} else {
		articleDetail, err := service.GetArticleDetailByID(aID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				ShowNotFound(c)
				return
			}

			ShowInternalServerError(c)
			return
		}

		renderData["Article"] = articleDetail
		renderData["LastArticle"] = service.GetLastArticle(aID)
		renderData["NextArticle"] = service.GetNextArticle(aID)
		renderData["Title"] = renderData["Article"].(*service.ShowArticleDetail).Title + " - " + renderData["Setting"].(map[string]interface{})["BlogName"].(string)

		// set cache
		articleDetailCache := map[string]interface{}{
			"Article":     renderData["Article"],
			"LastArticle": renderData["LastArticle"],
			"NextArticle": renderData["NextArticle"],
		}
		service.SrvEngine.MarshalAndSetCache(config.CacheArticleDetailPrefix+articleID, articleDetailCache)
	}

	counter.CounterCache.CountOne(c.ClientIP(), aID)

	c.HTML(http.StatusOK, getTheme(c)+"/article-detail.html", renderData)
}
