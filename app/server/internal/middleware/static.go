package middleware

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"path"
)

const (
	// StaticPath is the path to the static files
	staticPath = "static"
)

// Static returns a middleware handler that serves static files in the given directory.
func Static() gin.HandlerFunc {
	return static.Serve("/", static.LocalFile(staticPath, true))
}

// NoRoute SPA router
func NoRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.File(path.Join(staticPath, "index.html"))
	}
}
