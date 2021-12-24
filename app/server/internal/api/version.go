package api

import (
	"github.com/gin-gonic/gin"
	"github.com/iyear/pure-live/global"
	"github.com/iyear/pure-live/pkg/ecode"
	"github.com/iyear/pure-live/pkg/format"
)

func GetVersion(c *gin.Context) {
	format.HTTP(c, ecode.Success, nil, gin.H{
		"ver":     global.Version,
		"runtime": global.GetRuntime(),
	})
}
