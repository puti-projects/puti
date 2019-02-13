package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ShowPageDetail handle page info
func ShowPageDetail(c *gin.Context) {
	renderData := getRenderData(c)

	c.HTML(http.StatusOK, getTheme(c)+"/page-detail.html", renderData)
}
