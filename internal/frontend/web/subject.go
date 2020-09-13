package web

import (
	"errors"
	"net/http"

	"github.com/puti-projects/puti/internal/frontend/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ShowTopSubjects subject list (parent id is 0)
func ShowTopSubjects(c *gin.Context) {
	// get renderer data include basic data
	renderData := getRenderData(c)

	// get content
	subjectsList, err := service.GetChildrenSubejcts(0)
	if err != nil {
		ShowInternalServerError(c)
		return
	}
	renderData["SubjectList"] = subjectsList

	renderData["Widgets"] = getWidgets()
	renderData["Title"] = "专题" + " - " + renderData["Setting"].(map[string]interface{})["BlogName"].(string)
	c.HTML(http.StatusOK, getTheme(c)+"/subjects.html", renderData)
}

// ShowSubjects subejct list for children (parent id != 0)
func ShowSubjects(c *gin.Context) {
	renderData := getRenderData(c)

	// subject slug
	subjectSlug := c.Param("slug")

	// get subject info
	subjectInfo, err := service.GetSubjectInfoBySlug(subjectSlug)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ShowNotFound(c)
			return
		}

		ShowInternalServerError(c)
		return
	}

	// get subject children
	subjectsList, err := service.GetChildrenSubejcts(subjectInfo.ID)
	if err != nil {
		ShowInternalServerError(c)
		return
	}

	var subjectArticles []*map[string]interface{}
	// no children, then get article list
	if len(subjectsList) == 0 {
		subjectArticles, err = service.GetSubjectArticleList(subjectInfo.ID)
		if err != nil {
			ShowInternalServerError(c)
			return
		}
	}

	renderData["SubjectInfo"] = subjectInfo
	renderData["SubjectList"] = subjectsList
	renderData["SubjectArticles"] = subjectArticles

	renderData["Widgets"] = getWidgets()
	renderData["Title"] = subjectInfo.Name + " - 专题 - " + renderData["Setting"].(map[string]interface{})["BlogName"].(string)
	c.HTML(http.StatusOK, getTheme(c)+"/subjects.html", renderData)
}
