package bilibili

import (
	"encoding/binary"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/iyear/pure-live-core/model"
	"strconv"
)

func (c *base) Host(room string) string {
	_ = room
	return "wss://broadcastlv.chat.bilibili.com/sub"
}
func (c *base) Enter(room string) (int, [][]byte, error) {
	r, err := strconv.ParseInt(room, 10, 64)
	if err != nil {
		return 0, nil, err
	}

	enter := map[string]interface{}{
		"platform": "web",
		"protover": 2,
		"roomid":   r,
		"uid":      123456789,
		"type":     2,
		"key":      "",
	}
	body, err := json.Marshal(enter)
	if err != nil {
		return -1, nil, err
	}
	return websocket.BinaryMessage, [][]byte{encode(0, wsOpEnterRoom, body)}, nil
}
func (c *base) HeartBeat() (int, []byte, error) {
	return websocket.BinaryMessage, encode(0, wsOpHeartbeat, nil), nil
}
func (c *base) Handle(tp int, data []byte) ([]model.Msg, bool, error) {
	if tp != websocket.BinaryMessage || len(data) < 16 {
		return nil, false, nil
	}
	return c.handle(data)
}

func (c *base) handle(b []byte) ([]model.Msg, bool, error) {
	r := make([]model.Msg, 0)
	ver, op, body := decode(b)
	switch op {
	case wsOpHeartbeatReply:
		r = append(r, &model.MsgHot{Hot: int64(binary.BigEndian.Uint32(body))})
		return r, true, nil
	case wsOpMessage:
		// 压缩版本重新解包再调用，直到 ver==0
		switch ver {
		case wsVerPlain:
			return c.handlePlain(body)
		case wsVerZlib:
			de, err := zlibDe(body)
			if err != nil {
				return nil, false, err
			}
			return c.handles(c.split(de))
		case wsVerBrotli:
			de, err := brotliDe(body)
			if err != nil {
				return nil, false, err
			}
			return c.handles(c.split(de))
		}
	}
	return nil, false, nil
}

// split 压缩过的body需要拆包
func (c *base) split(b []byte) [][]byte {
	var packs [][]byte
	for i, size := uint32(0), uint32(0); i < uint32(len(b)); i += size {
		size = binary.BigEndian.Uint32(b[i : i+4])
		packs = append(packs, b[i:i+size])
	}
	return packs
}
func (c *base) handles(bs [][]byte) ([]model.Msg, bool, error) {
	r := make([]model.Msg, 0)
	for _, b := range bs {
		m, ok, err := c.handle(b)
		if err != nil {
			continue
		}
		if !ok {
			continue
		}
		r = append(r, m...)
	}
	return r, true, nil
}
func (c *base) handlePlain(body []byte) ([]model.Msg, bool, error) {
	var cmd struct {
		CMD string `json:"cmd"`
	}

	if err := json.Unmarshal(body, &cmd); err != nil {
		return nil, false, err
	}
	m, ok, err := c.switchCmd(cmd.CMD, body)
	return []model.Msg{m}, ok, err
}
func (c *base) switchCmd(cmd string, body []byte) (model.Msg, bool, error) {
	switch cmd {
	case "DANMU_MSG":
		d, err := parseDanmaku(body)
		if err != nil {
			return nil, false, err
		}
		return &model.MsgDanmaku{
			Content: d.Content,
			Type:    bilibili2dp(d.SendMode),
			Color:   d.DanmakuColor,
		}, true, nil
	}
	return nil, false, nil
}

type Danmaku struct {
	SendMode     int    `json:"send_mode"`
	SendFontSize int    `json:"send_font_size"`
	DanmakuColor int64  `json:"danmaku_color"`
	Time         int64  `json:"time"`
	DMID         int64  `json:"dmid"`
	MsgType      int    `json:"msg_type"`
	Bubble       string `json:"bubble"`
	Content      string `json:"content"`
	MID          int64  `json:"mid"`
	Uname        string `json:"uname"`
	RoomAdmin    int    `json:"room_admin"`
	Vip          int    `json:"vip"`
	SVip         int    `json:"svip"`
	Rank         int    `json:"rank"`
	MobileVerify int    `json:"mobile_verify"`
	UnameColor   string `json:"uname_color"`
	MedalName    string `json:"medal_name"`
	UpName       string `json:"up_name"`
	MedalLevel   int    `json:"medal_level"`
	UserLevel    int    `json:"user_level"`
}

func parseDanmaku(b []byte) (*Danmaku, error) {
	var t map[string]interface{}
	if err := json.Unmarshal(b, &t); err != nil {
		return nil, err
	}
	info := t["info"]
	var dm = &Danmaku{}
	l := len(info.([]interface{}))
	if l >= 1 {
		h := info.([]interface{})[0].([]interface{})
		dm.SendMode = int(h[1].(float64))
		dm.SendFontSize = int(h[2].(float64))
		dm.DanmakuColor = int64(h[3].(float64))
		dm.Time = int64(h[4].(float64))
		dm.DMID = int64(h[5].(float64))
		dm.MsgType = int(h[10].(float64))
		dm.Bubble = h[11].(string)
	}
	if l >= 2 {
		dm.Content = info.([]interface{})[1].(string)
	}
	if l >= 3 {
		h := info.([]interface{})[2].([]interface{})
		dm.MID = int64(h[0].(float64))
		dm.Uname = h[1].(string)
		dm.RoomAdmin = int(h[2].(float64))
		dm.Vip = int(h[3].(float64))
		dm.SVip = int(h[4].(float64))
		dm.Rank = int(h[5].(float64))
		dm.MobileVerify = int(h[6].(float64))
		dm.UnameColor = h[7].(string)
	}
	if l >= 4 {
		h := info.([]interface{})[3].([]interface{})
		l2 := len(h)
		if l2 >= 1 {
			dm.MedalLevel = int(h[0].(float64))
		}
		if l2 >= 2 {
			dm.MedalName = h[1].(string)
		}
		if l2 >= 3 {
			dm.UpName = h[2].(string)
		}
	}
	if l >= 5 {
		dm.UserLevel = int(info.([]interface{})[4].([]interface{})[0].(float64))
	}
	return dm, nil
}
