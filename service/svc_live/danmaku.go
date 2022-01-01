package svc_live

import (
	"github.com/iyear/pure-live-core/model"
)

func SendDanmaku(client model.Client, room, content string, tp int, color int64) error {
	return client.SendDanmaku(room, content, tp, color)
}
