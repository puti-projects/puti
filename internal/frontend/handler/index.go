package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowIndex(c *gin.Context) {
	renderData := getRenderData(c)

	c.HTML(http.StatusOK, getTheme(c)+"/index.html", renderData)
}
