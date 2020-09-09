package utils

import (
	"github.com/gin-gonic/gin"
)

// GetReqID get X-Request-Id header and return the value which is a uuid
func GetReqID(c *gin.Context) string {
	v, ok := c.Get("X-Request-Id")
	if !ok {
		return ""
	}

	if requestID, ok := v.(string); ok {
		return requestID
	}
	return ""
}
