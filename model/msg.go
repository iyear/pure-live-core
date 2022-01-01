package model

import "github.com/iyear/pure-live-core/pkg/conf"

type Msg interface {
	Event() string
}

type MsgDanmaku struct {
	Content string `json:"content"`
	Type    int    `json:"type"`
	Color   int64  `json:"color"`
}

func (m *MsgDanmaku) Event() string {
	return conf.EventDanmaku
}

type MsgHot struct {
	Hot int64 `json:"hot"`
}

func (m *MsgHot) Event() string {
	return conf.EventHot
}
