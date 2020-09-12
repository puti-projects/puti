package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/puti-projects/puti/internal/frontend/service"
)

// ShowArchive handle archive list
func ShowArchive(c *gin.Context) {
	// get renderer data include basic data
	renderData := getRenderData(c)

	// get content
	archive, sortYear, sortMonth, err := service.GetArchive()
	if err != nil {
		ShowInternalServerError(c)
		return
	}

	renderData["Archive"] = archive
	renderData["ArchiveSortYear"] = sortYear
	renderData["ArchiveSortMonth"] = sortMonth

	renderData["Widgets"] = getWidgets()
	renderData["Title"] = "归档" + " - " + renderData["Setting"].(map[string]interface{})["BlogName"].(string)
	c.HTML(http.StatusOK, getTheme(c)+"/archive.html", renderData)
}
