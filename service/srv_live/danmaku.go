package srv_live

import (
	"github.com/iyear/pure-live/global"
)

func SendDanmaku(id string, content string, tp int, color int64) error {
	var (
		conn *global.Conn
		err  error
	)
	if conn, err = global.GetConn(id); err != nil {
		return err
	}
	if err = conn.Client.SendDanmaku(conn.Room, content, tp, color); err != nil {
		return err
	}
	return nil
}
