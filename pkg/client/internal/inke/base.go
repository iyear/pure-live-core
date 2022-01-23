package inke

import (
	"github.com/iyear/pure-live-core/model"
	"github.com/iyear/pure-live-core/pkg/client/internal/abstract"
	"github.com/iyear/pure-live-core/pkg/conf"
	"github.com/iyear/pure-live-core/pkg/util"
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
