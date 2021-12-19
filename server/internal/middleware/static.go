package middleware

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func Static() gin.HandlerFunc {
	return static.Serve("/", static.LocalFile("./static", false))
}
