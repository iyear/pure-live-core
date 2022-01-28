package inke

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/iyear/pure-live-core/model"
	"github.com/iyear/pure-live-core/pkg/client/internal/abstract"
	"github.com/iyear/pure-live-core/pkg/conf"
	"github.com/iyear/pure-live-core/pkg/util"
	"github.com/tidwall/gjson"
	"strconv"
)

type Inke struct {
	*abstract.Client
}

func NewInke() (model.Client, error) {
	return &Inke{}, nil
}

func (i *Inke) Plat() string {
	return conf.PlatInke
}

func (i *Inke) GetPlayURL(room string, qn int) (*model.PlayURL, error) {
	info, err := getRoomInfo(room)
	if err != nil {
		return nil, err
	}

	return &model.PlayURL{
		Qn:     qn,
		Desc:   util.Qn2Desc(qn),
		Origin: info.Get("data.live_addr.0.stream_addr").String(),
		CORS:   false,
		Type:   conf.StreamFlv,
	}, nil
}

// GetRoomInfo inke对于普通用户和未开播的主播返回的都是"当前用户无直播"，所以只能都返回未知主播
func (i *Inke) GetRoomInfo(room string) (*model.RoomInfo, error) {
	info, err := getRoomInfo(room)
	if err != nil {
		return &model.RoomInfo{
			Status: 0,
			Room:   room,
			Upper:  "未知主播",
			Link:   getRoomLink(room),
			Title:  "",
		}, nil
	}

	return &model.RoomInfo{
		Status: 1,
		Room:   room,
		Upper:  info.Get("data.media_info.nick").String(),
		Link:   getRoomLink(room),
		Title:  info.Get("data.live_name").String(),
	}, nil
}

func (i *Inke) Host(room string) string {
	info, err := getRoomInfo(room)
	if err != nil {
		return ""
	}
	// fmt.Println(info.Get("data.sio_url").String())
	return info.Get("data.sio_url").String()
}

// HeartBeat 无需心跳
func (i *Inke) HeartBeat() (int, []byte, error) {
	return 0, nil, conf.ErrSkip
}

// Enter 无需进入信息
func (i *Inke) Enter(room string) (int, [][]byte, error) {
	_ = room
	return 0, nil, conf.ErrSkip
}
func (i *Inke) Handle(tp int, data []byte) ([]model.Msg, bool, error) {
	if tp != websocket.TextMessage {
		return nil, false, nil
	}
	j := gjson.Parse(string(data))

	if j.Get("b.ev").String() != "s.m" {
		return nil, false, nil
	}

	msgs := make([]model.Msg, 0)
	// 还有一种 color type，消息是xxx来到直播间，这里忽略该种消息，只做弹幕
	for _, msg := range j.Get("ms").Array() {
		if msg.Get("tp").String() == "pub" {
			fmt.Println()
			for _, t := range msg.Get("statement_info").Array() {
				color, err := strconv.ParseInt(t.Get("c_color").String()[1:], 16, 64)
				if err != nil {
					continue
				}
				msgs = append(msgs, &model.MsgDanmaku{
					Content: t.Get("c").String(),
					Type:    conf.DanmakuTypeRight,
					Color:   color,
				})
			}
		}
	}
	return msgs, true, nil
}
