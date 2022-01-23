package svc_live

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/iyear/pure-live-core/model"
	"github.com/iyear/pure-live-core/pkg/conf"
	"go.uber.org/zap"
	"time"
)

func Serve(ctx context.Context, dialer *websocket.Dialer, id string, client model.Client, room string) (chan *model.Transport, error) {
	// 过程
	// conn -> enter(on entered) -> go receive()   -> for{send msg to local}
	//  						 -> go heartbeat() -

	live, _, err := dialer.DialContext(ctx, client.Host(room), nil)
	if err != nil {
		return nil, err
	}

	zap.S().Infow("connected to live danmaku server", "id", id)

	rev := make(chan *model.Transport)

	tp, data, err := client.Enter(room)

	if err != nil && err != conf.ErrSkip {
		return nil, err
	}
	for _, d := range data {
		if err = live.WriteMessage(tp, d); err != nil {
			return nil, err
		}
	}

	zap.S().Infow("entered the room", "id", id)

	go receive(ctx, id, client, live, rev)
	go heartbeat(ctx, id, client, live)

	return rev, nil
}

func heartbeat(ctx context.Context, id string, client model.Client, live *websocket.Conn) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	hb := func() {
		tp, data, err := client.HeartBeat()
		if err == conf.ErrSkip {
			return
		}
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

func receive(ctx context.Context, id string, client model.Client, live *websocket.Conn, rev chan *model.Transport) {
	msgCtx, msgCancel := context.WithCancel(ctx)
	defer msgCancel()
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
			msg, ok, err := client.Handle(t, data)

			// 错误判断
			if err != nil {
				go push(msgCtx, id, &model.Transport{
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
				go push(msgCtx, id, &model.Transport{
					Msg:   m,
					Error: nil,
				}, rev)
			}
		}
	}
}
func push(ctx context.Context, id string, tp *model.Transport, rev chan *model.Transport) {
	t := time.NewTimer(5 * time.Second)
	defer t.Stop()

	select {
	case <-t.C:
	case <-ctx.Done():
		zap.S().Infow("push stopped", "id", id)
	case rev <- tp:
	}
}
