package middleware

import (
	"time"

	"github.com/puti-projects/puti/internal/pkg/logger"
	"github.com/puti-projects/puti/internal/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// AccessLogger replace gin default access logger
// Note: using zap logger
// In puti, runmode in release will not generate log infomation
// In production, it is better to let web server (nginx) to do this job
func AccessLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		// Process requests
		c.Next()

		// log part
		cost := time.Since(start)
		logger.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
			zap.String("request-ID", utils.GetReqID(c)),
		)
	}
}
