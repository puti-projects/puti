package middleware

import (
	"github.com/puti-projects/puti/internal/backend/handler"
	"github.com/puti-projects/puti/internal/pkg/errno"
	"github.com/puti-projects/puti/internal/pkg/token"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware middleware of token prase
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Prase the json web token
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
