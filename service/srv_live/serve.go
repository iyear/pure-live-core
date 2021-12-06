package srv_live

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/iyear/pure-live/api"
	"github.com/iyear/pure-live/conf"
	"github.com/iyear/pure-live/global"
	"github.com/iyear/pure-live/model"
	"github.com/iyear/pure-live/util"
	"go.uber.org/zap"
	"time"
)

func Serve(ctx context.Context) {
	// 过程
	// conn -> enter(on entered) -> go receive()   -> for{send msg to local}
	//  						 -> go heartbeat() -

	id := ctx.Value("id").(string)
	conn, err := global.GetConn(id)
	if err != nil {
		return
	}

	dialer := websocket.DefaultDialer
	if conf.C.Socks5.Enable {
		dialer.NetDial = util.MustGetSocks5(conf.C.Socks5.Host, conf.C.Socks5.Port, conf.C.Socks5.User, conf.C.Socks5.Password).Dial
	}

	live, _, err := dialer.DialContext(ctx, conn.Client.Host(), nil)
	if err != nil {
		zap.S().Warnw("failed to connect live websocket server", "id", id, "error", err)
		return
	}
	defer live.Close()

	zap.S().Infow("connected to live danmaku server", "id", id)

	rev := make(chan *model.Transport)

	tp, data, err := conn.Client.Enter(conn.Room)
	// fmt.Println(hex.Dump(data))
	if err != nil {
		zap.S().Warnw("failed to get enter room data", "id", id, "error", err)
		return
	}
	for _, d := range data {
		if err = live.WriteMessage(tp, d); err != nil {
			zap.S().Warnw("failed to write ws msg", "id", id, "error", err)
			return
		}
	}

	zap.S().Infow("entered the room", "id", id)

	hbCtx, hbCancel := context.WithCancel(ctx)
	revCtx, revCancel := context.WithCancel(ctx)

	go receive(revCtx, live, rev)
	defer revCancel()

	go heartbeat(hbCtx, live)
	defer hbCancel()

	// 客户端心跳检查
	health := time.NewTicker(5 * time.Second)
	defer health.Stop()
	for {
		select {
		// 5秒检查一次客户端存活
		case <-health.C:
			if err = conn.Server.WriteMessage(websocket.TextMessage, api.MsgFmt(conf.EventCheck, nil)); err != nil {
				zap.S().Warnw("failed to write ws message", "id", id, "error", err)
				return
			}
		case tp := <-rev:
			if tp.Error != nil {
				zap.S().Warnw("receive transport error", "id", id, "error", err)
				continue
			}
			if err = conn.Server.WriteMessage(websocket.TextMessage, api.MsgFmt(tp.Msg.Event(), tp.Msg)); err != nil {
				zap.S().Warnw("failed to write ws message", "id", id, "error", err)
				return
			}
		}
	}
}

func heartbeat(ctx context.Context, live *websocket.Conn) {
	id := ctx.Value("id").(string)
	conn, err := global.GetConn(id)
	if err != nil {
		return
	}

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	hb := func() {
		tp, data, err := conn.Client.HeartBeat()
		if err != nil {
			zap.S().Warnw("failed to get heartbeat data", "id", id, "error", err)
			return
		}
		if err = live.WriteMessage(tp, data); err != nil {
			zap.S().Warnw("failed to send heartbeat", "id", id, "error", err)
			return
		}
	}
	// 开头先执行一次
	hb()
	for {
		select {
		case <-ctx.Done():
			zap.S().Infow("heartbeat stopped",
				"id", id)
			return
		case <-ticker.C:
			hb()
		}
	}

}
func receive(ctx context.Context, live *websocket.Conn, rev chan *model.Transport) {
	msgCtx, msgCancel := context.WithCancel(ctx)
	defer msgCancel()

	var (
		conn *global.Conn
		err  error
	)
	id := ctx.Value("id").(string)
	if conn, err = global.GetConn(id); err != nil {
		return
	}

	for {
		select {
		case <-ctx.Done():
			zap.S().Infow("receive stopped", "id", id)
			return
		default:
			t, data, err := live.ReadMessage()
			if err != nil {
				continue
			}
			msg, ok, err := conn.Client.Handle(t, data)

			// 错误判断
			if err != nil {
				go push(msgCtx, &model.Transport{
					Msg:   nil,
					Error: err,
				}, rev)
				continue
			}
			// 跳过判断
			if !ok {
				continue
			}

			for _, m := range msg {
				go push(msgCtx, &model.Transport{
					Msg:   m,
					Error: nil,
				}, rev)
			}
		}
	}
}
func push(ctx context.Context, tp *model.Transport, rev chan *model.Transport) {
	id := ctx.Value("id").(string)

	t := time.NewTimer(5 * time.Second)
	defer t.Stop()

	select {
	case <-t.C:
	case <-ctx.Done():
		zap.S().Infow("push stopped", "id", id)
	case rev <- tp:
	}
}
