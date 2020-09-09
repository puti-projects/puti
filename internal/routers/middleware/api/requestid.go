package api

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/google/uuid"
)

// RequestID middleware of set a X-Request-Id header by general a uuid
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// check for incoming header, use it if exists
		requestID := c.Request.Header.Get("X-Request-Id")

		// Create request id with UUID4
		if requestID == "" {
			u4, _ := uuid.NewRandom()
			requestID = u4.String()
		}

		// Expose it for use in the application
		c.Set("X-Request-Id", requestID)

		// Set X-Request-Id header
		c.Writer.Header().Set("X-Request-Id", requestID)
		c.Next()
	}
}
