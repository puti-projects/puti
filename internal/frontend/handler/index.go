package handler

import (
	"net/http"

	optionCache "github.com/puti-projects/puti/internal/pkg/option"

	"github.com/gin-gonic/gin"
)

func ShowIndex(c *gin.Context) {

	c.HTML(http.StatusOK, optionCache.Options.Get("current_theme")+"/index.html", gin.H{})
}
