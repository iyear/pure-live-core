package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/iyear/pure-live/global"
	"github.com/iyear/pure-live/pkg/client"
	"github.com/iyear/pure-live/pkg/e"
	"github.com/iyear/pure-live/server/api"
	"github.com/iyear/pure-live/service/srv_live"
	"go.uber.org/zap"
	"net/http"
)

func Serve(c *gin.Context) {
	var err error
	req := struct {
		Plat string `form:"plat" binding:"required,max=15" json:"plat"`
		Room string `form:"room" binding:"required" json:"room"`
	}{}
	if err = c.ShouldBind(&req); err != nil {
		api.RespFmt(c, e.InvalidParams, err, nil)
		return
	}

	up := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	id := ""
	for {
		id = uuid.New().String()
		if _, ok := global.Hub.Conn.Load(id); !ok {
			break
		}
	}

	header := http.Header{}
	cookie := http.Cookie{
		Name:     "uuid",
		Path:     "/",
		Value:    id,
		Secure:   false,
		HttpOnly: false,
	}
	header.Set("Set-Cookie", cookie.String())

	cli, err := client.GetClient(req.Plat)
	if err != nil {
		zap.S().Warnw("failed to get platform", "id", id, "error", err, "plat", req.Plat)
		api.RespFmt(c, e.UnknownError, err, nil)
		return
	}
	defer cli.Stop()

	ws := &websocket.Conn{}
	if ws, err = up.Upgrade(c.Writer, c.Request, header); err != nil {
		zap.S().Errorw("failed to upgrade to websocket connection", "id", id, "error", err)
		return
	}
	defer func(ws *websocket.Conn) {
		_ = ws.Close()
	}(ws)

	global.Hub.Conn.Store(id, &global.Conn{
		Server: ws,
		Room:   req.Room,
		Client: cli,
	})
	defer global.Hub.Conn.Delete(id)

	ctx, stop := context.WithCancel(context.WithValue(context.Background(), "id", id))
	defer stop()

	zap.S().Infow("start serving...", "id", id, "room", req.Room, "plat", req.Plat)

	srv_live.Serve(ctx)

	zap.S().Infow("stop serving...", "id", id, "room", req.Room, "plat", req.Plat)
}

func GetPlayURL(c *gin.Context) {
	req := struct {
		Plat string `form:"plat" binding:"required,max=15" json:"plat"`
		Room string `form:"room" binding:"required" json:"room"`
	}{}
	if err := c.ShouldBind(&req); err != nil {
		api.RespFmt(c, e.InvalidParams, nil, nil)
		return
	}
	url, err := srv_live.GetPlayURL(req.Plat, req.Room)
	if err != nil {
		api.RespFmt(c, e.ErrorGetPlayURL, err, nil)
		zap.S().Warnw("failed to get play url", "error", err, "req", req)
		return
	}
	api.RespFmt(c, e.Success, nil, url)
}
func GetRoomInfo(c *gin.Context) {
	req := struct {
		Plat string `form:"plat" binding:"required,max=15" json:"plat"`
		Room string `form:"room" binding:"required" json:"room"`
	}{}
	if err := c.ShouldBind(&req); err != nil {
		api.RespFmt(c, e.InvalidParams, nil, nil)
		return
	}
	info, err := srv_live.GetRoomInfo(req.Plat, req.Room)
	if err != nil {
		api.RespFmt(c, e.ErrorGetRoomInfo, err, nil)
		zap.S().Warnw("failed to get room info", "error", err, "req", req)
		return
	}
	api.RespFmt(c, e.Success, nil, info)
}
func SendDanmaku(c *gin.Context) {
	req := struct {
		ID      string `form:"id" binding:"required,uuid"` // 服务端分发的uuid
		Content string `form:"content" binding:"required" json:"content"`
		Type    int    `form:"type" binding:"required,gte=0,lte=2" json:"type"` // 1:顶部 0:滚动 2:底部
		Color   int64  `form:"color" binding:"required" json:"color"`
	}{}
	if err := c.ShouldBind(&req); err != nil {
		api.RespFmt(c, e.InvalidParams, nil, nil)
		return
	}
	if err := srv_live.SendDanmaku(req.ID, req.Content, req.Type, req.Color); err != nil {
		zap.S().Warnw("failed to send danmaku", "error", err, "req", req)
		api.RespFmt(c, e.ErrorSendDanmaku, err, nil)
		return
	}
	api.RespFmt(c, e.Success, nil, nil)
}
