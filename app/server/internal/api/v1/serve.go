package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/iyear/pure-live-core/app/server/internal/config"
	"github.com/iyear/pure-live-core/global"
	"github.com/iyear/pure-live-core/pkg/client"
	"github.com/iyear/pure-live-core/pkg/conf"
	"github.com/iyear/pure-live-core/pkg/ecode"
	"github.com/iyear/pure-live-core/pkg/format"
	"github.com/iyear/pure-live-core/pkg/util"
	"github.com/iyear/pure-live-core/service/svc_live"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func Serve(c *gin.Context) {
	req := struct {
		Plat string `form:"plat" binding:"required,max=15" json:"plat"`
		Room string `form:"room" binding:"required" json:"room"`
	}{}
	if err := c.ShouldBind(&req); err != nil {
		format.HTTP(c, ecode.InvalidParams, err, nil)
		return
	}

	upgrader := websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
	id := getUniqueID()

	cli, err := client.GetClient(req.Plat)
	if err != nil {
		zap.S().Warnw("failed to get platform", "id", id, "error", err, "plat", req.Plat)
		format.HTTP(c, ecode.UnknownError, err, nil)
		return
	}
	defer cli.Stop()

	srvWS := &websocket.Conn{}
	if srvWS, err = upgrader.Upgrade(c.Writer, c.Request, getUpgradeHeader(id)); err != nil {
		zap.S().Errorw("failed to upgrade to websocket connection", "id", id, "error", err)
		return
	}
	defer func(ws *websocket.Conn) {
		_ = ws.Close()
	}(srvWS)

	global.Hub.Conn.Store(id, &global.Conn{
		Server: srvWS,
		Room:   req.Room,
		Client: cli,
	})
	defer global.Hub.Conn.Delete(id)

	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	zap.S().Infow("start serving...", "id", id, "room", req.Room, "plat", req.Plat)

	dialer := websocket.DefaultDialer
	if config.Server.Socks5.Enable {
		dialer.NetDial = util.MustGetSocks5(config.Server.Socks5.Host, config.Server.Socks5.Port, config.Server.Socks5.User, config.Server.Socks5.Password).Dial
	}

	rev, err := svc_live.Serve(ctx, dialer, id, cli, req.Room)
	if err != nil {
		zap.S().Warnw("failed to start serve", "id", id, "error", err)
		return
	}

	// 开始获取数据
	// 客户端心跳检查
	health := time.NewTicker(5 * time.Second)
	defer health.Stop()
	defer zap.S().Infow("stop serving...", "id", id, "room", req.Room, "plat", req.Plat)

	for {
		select {
		// 5秒检查一次客户端存活
		case <-health.C:
			if err = srvWS.WriteMessage(websocket.TextMessage, format.WS(conf.EventCheck, nil)); err != nil {
				zap.S().Warnw("failed to write ws message", "id", id, "error", err)
				return
			}
		case transport := <-rev:
			if transport.Error != nil {
				zap.S().Warnw("receive transport error", "id", id, "error", err)
				continue
			}
			if err = srvWS.WriteMessage(websocket.TextMessage, format.WS(transport.Msg.Event(), transport.Msg)); err != nil {
				zap.S().Warnw("failed to write ws message", "id", id, "error", err)
				return
			}
		}
	}
}

func getUpgradeHeader(id string) http.Header {
	header := http.Header{}
	cookie := http.Cookie{
		Name:     "uuid",
		Path:     "/",
		Value:    id,
		Secure:   false,
		HttpOnly: false,
	}
	header.Set("Set-Cookie", cookie.String())
	return header
}

func getUniqueID() string {
	id := ""
	for {
		id = util.RandLetters(10)
		if _, ok := global.Hub.Conn.Load(id); !ok {
			break
		}
	}
	return id
}
