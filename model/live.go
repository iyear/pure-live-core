package model

import "fmt"

// PlayURL live stream info
type PlayURL struct {
	Qn     int    `json:"qn"`
	Desc   string `json:"desc"`
	Origin string `json:"origin"`
	CORS   bool   `json:"cors"`
	Type   string `json:"type"`
}

// RoomInfo live room info
type RoomInfo struct {
	Status int    `json:"status"` // 0:未开播 1:已开播
	Room   string `json:"room"`   // 真实房间号
	Upper  string `json:"upper"`  // 主播名称
	Link   string `json:"link"`   // 直播间地址
	Title  string `json:"title"`  // 直播间标题
}

func (p *PlayURL) String() string {
	return fmt.Sprintf("qn: %d,desc: %s,type: %s\nurl: %s", p.Qn, p.Desc, p.Type, p.Origin)
}
func (r *RoomInfo) String() string {
	return fmt.Sprintf("status: %d,upper: %s,link: %s,room: %s,title: %s", r.Status, r.Upper, r.Link, r.Room, r.Title)
}
