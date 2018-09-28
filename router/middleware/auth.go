package middleware

import (
	"puti/handler"
	"puti/pkg/errno"
	"puti/pkg/token"

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
