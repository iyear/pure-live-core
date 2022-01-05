package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		zap.S().Infof("%3d | %13v | %15s | %s | %s |",
			c.Writer.Status(),
			time.Now().Sub(start),
			c.ClientIP(),
			c.Request.Method,
			c.Request.RequestURI)
	}
}
