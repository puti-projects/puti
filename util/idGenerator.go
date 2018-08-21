package util

import (
	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
)

// GenShortID ...
func GenShortID() (string, error) {
	return shortid.Generate()
}

// GetReqID ...
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
