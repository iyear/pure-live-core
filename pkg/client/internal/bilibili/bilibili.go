package bilibili

import (
	"fmt"
	"github.com/iyear/biligo"
	"strconv"
)

type BiliBili struct {
	client *biligo.BiliClient
	base
}
type BiliComm struct {
	client *biligo.CommClient
	base
}

func (c *BiliBili) SendDanmaku(room string, content string, tp int, color int64) error {
	var (
		roomm int64
		err   error
	)
	if roomm, err = strconv.ParseInt(room, 10, 64); err != nil {
		return err
	}
	if err = c.client.LiveSendDanmaku(roomm, color, 25, dp2bilibili(tp), content, 0); err != nil {
		return err
	}
	return nil
}

func (c *BiliComm) SendDanmaku(room string, content string, tp int, color int64) error {
	_ = room
	_ = content
	_ = tp
	_ = color
	return fmt.Errorf("this type of client does not support sending danmaku")
}
