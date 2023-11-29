package web

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// RedirectToWWW redirect to www
func RedirectToWWW(c *gin.Context) {
	host := c.Request.Host
	// 如果不是本地域名
	if !strings.Contains(host, "localhost") && !strings.Contains(host, "127.0.0.1") {
		if !strings.HasPrefix(host, "www.") {
			c.Redirect(http.StatusMovedPermanently, "https://www."+host+c.Request.RequestURI)
			c.Abort()
		}
	}
}
