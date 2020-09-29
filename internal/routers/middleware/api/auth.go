package api

import (
	"github.com/puti-projects/puti/internal/admin/api"
	"github.com/puti-projects/puti/internal/pkg/errno"
	"github.com/puti-projects/puti/internal/pkg/token"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware middleware of token prase
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Prase the json view token
		if _, err := token.ParseRequest(c); err != nil {
			api.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
