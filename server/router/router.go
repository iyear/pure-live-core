package router

import (
	"github.com/gin-gonic/gin"
	"github.com/iyear/pure-live/pkg/conf"
	"github.com/iyear/pure-live/pkg/util"
	v12 "github.com/iyear/pure-live/server/api/v1"
	middleware2 "github.com/iyear/pure-live/server/middleware"
)

var r *gin.Engine

func Init() *gin.Engine {
	gin.SetMode(util.IF(conf.C.Server.Debug, gin.DebugMode, gin.ReleaseMode).(string))
	r = gin.New()

	r.Use(middleware2.Recovery())
	r.Use(middleware2.CORS())
	r.Use(middleware2.Static())

	g := r.Group("/api")
	apiV1 := g.Group("/v1")
	{
		apiV1.GET("/live/serve", v12.Serve)
		apiV1.GET("/live/play", v12.Play)
		apiV1.GET("/live/room_info", v12.GetRoomInfo)
		apiV1.GET("/live/play_url", v12.GetPlayURL)
		apiV1.POST("/live/danmaku/send", v12.SendDanmaku)

		apiV1.POST("/fav/list/add", v12.AddFavList)
		apiV1.GET("/fav/list/get_all", v12.GetAllFavLists)
		apiV1.POST("/fav/list/del", v12.DelFavList)
		apiV1.POST("/fav/list/edit", v12.EditFavList)
		apiV1.GET("/fav/list/get", v12.GetFavList)

		apiV1.GET("/fav/get", v12.GetFav)
		apiV1.POST("/fav/add", v12.AddFav)
		apiV1.POST("/fav/del", v12.DelFav)
		apiV1.POST("/fav/edit", v12.EditFav)
	}

	return r
}
