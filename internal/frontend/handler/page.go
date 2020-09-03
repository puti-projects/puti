package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/puti-projects/puti/internal/frontend/service"
	"github.com/puti-projects/puti/internal/pkg/counter"
)

// ShowPageDetail handle page info
// TODO get page template setting
func ShowPageDetail(c *gin.Context, pageID uint64) {
	renderData := getRenderData(c)

	pageDetail, err := service.GetPageDetailByID(pageID)
	if err != nil {
		ShowInternalServerError(c)
		return
	}

	counter.CounterCache.CountOne(c.ClientIP(), pageID)

	renderData["Page"] = pageDetail

	renderData["Widgets"] = getWidgets()
	renderData["Title"] = pageDetail.Title + " - " + renderData["Setting"].(map[string]interface{})["BlogName"].(string)
	c.HTML(http.StatusOK, getTheme(c)+"/page-detail.html", renderData)
}
