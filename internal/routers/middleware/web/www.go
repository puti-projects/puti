package web

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// RedirectToWWW redirect to www
func RedirectToWWW(c *gin.Context) {
	host := c.Request.Host
	if !strings.HasPrefix(host, "www.") {
		c.Redirect(http.StatusMovedPermanently, "https://www."+host+c.Request.RequestURI)
		c.Abort()
	}
}
