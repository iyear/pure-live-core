package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/iyear/pure-live/pkg/ecode"
	"github.com/iyear/pure-live/pkg/format"
	"github.com/iyear/pure-live/service/srv_live"
	"go.uber.org/zap"
)

func GetPlayURL(c *gin.Context) {
	req := struct {
		Plat string `form:"plat" binding:"required,max=15" json:"plat"`
		Room string `form:"room" binding:"required" json:"room"`
	}{}
	if err := c.ShouldBind(&req); err != nil {
		format.HTTP(c, ecode.InvalidParams, nil, nil)
		return
	}
	url, err := srv_live.GetPlayURL(req.Plat, req.Room)
	if err != nil {
		format.HTTP(c, ecode.ErrorGetPlayURL, err, nil)
		zap.S().Warnw("failed to get play url", "error", err, "req", req)
		return
	}
	format.HTTP(c, ecode.Success, nil, url)
}
func GetRoomInfo(c *gin.Context) {
	req := struct {
		Plat string `form:"plat" binding:"required,max=15" json:"plat"`
		Room string `form:"room" binding:"required" json:"room"`
	}{}
	if err := c.ShouldBind(&req); err != nil {
		format.HTTP(c, ecode.InvalidParams, nil, nil)
		return
	}
	info, err := srv_live.GetRoomInfo(req.Plat, req.Room)
	if err != nil {
		format.HTTP(c, ecode.ErrorGetRoomInfo, err, nil)
		zap.S().Warnw("failed to get room info", "error", err, "req", req)
		return
	}
	format.HTTP(c, ecode.Success, nil, info)
}
func SendDanmaku(c *gin.Context) {
	req := struct {
		ID      string `form:"id" binding:"required,uuid"` // 服务端分发的uuid
		Content string `form:"content" binding:"required" json:"content"`
		Type    int    `form:"type" binding:"required,gte=0,lte=2" json:"type"` // 1:顶部 0:滚动 2:底部
		Color   int64  `form:"color" binding:"required" json:"color"`
	}{}
	if err := c.ShouldBind(&req); err != nil {
		format.HTTP(c, ecode.InvalidParams, nil, nil)
		return
	}
	if err := srv_live.SendDanmaku(req.ID, req.Content, req.Type, req.Color); err != nil {
		zap.S().Warnw("failed to send danmaku", "error", err, "req", req)
		format.HTTP(c, ecode.ErrorSendDanmaku, err, nil)
		return
	}
	format.HTTP(c, ecode.Success, nil, nil)
}
