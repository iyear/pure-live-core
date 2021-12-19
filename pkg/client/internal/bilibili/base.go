package bilibili

import (
	"github.com/iyear/biligo"
	"github.com/iyear/pure-live/model"
	"github.com/iyear/pure-live/pkg/conf"
	"github.com/iyear/pure-live/pkg/util"
	"strconv"
)

type base struct{}

func NewBiliBili() (model.Client, error) {
	if !conf.C.Account.BiliBili.Enable {
		return &BiliComm{
			client: biligo.NewCommClient(&biligo.CommSetting{
				DebugMode: conf.C.Server.Debug,
			}),
		}, nil
	}
	b, err := biligo.NewBiliClient(&biligo.BiliSetting{
		Auth: &biligo.CookieAuth{
			DedeUserID:      conf.C.Account.BiliBili.DedeUserID,
			DedeUserIDCkMd5: conf.C.Account.BiliBili.DedeUserIDCkMd5,
			SESSDATA:        conf.C.Account.BiliBili.SESSDATA,
			BiliJCT:         conf.C.Account.BiliBili.BiliJCT,
		},
		DebugMode: conf.C.Server.Debug,
	})
	if err != nil {
		return nil, err
	}
	return &BiliBili{client: b}, nil
}

func (c *base) Plat() string {
	return conf.PlatBiliBili
}

func (c *base) GetPlayURL(room string, qn int) (*model.PlayURL, error) {
	client := biligo.NewCommClient(&biligo.CommSetting{})
	roomNum, err := strconv.ParseInt(room, 10, 64)
	if err != nil {
		return nil, err
	}

	// 内部维护一个qn映射表
	q := map[int]int{
		conf.QnBest: 10000,
		conf.QnHigh: 400,
		conf.QnMid:  250,
		conf.QnLow:  80,
	}

	r, err := client.LiveGetPlayURL(roomNum, q[qn])
	if err != nil {
		return nil, err
	}
	return &model.PlayURL{
		Qn:     qn,
		Desc:   util.Qn2Desc(qn),
		Origin: r.DURL[0].URL,
		CORS:   false,
		Type:   conf.StreamFlv,
	}, nil
}

func (c *base) GetRoomInfo(room string) (*model.RoomInfo, error) {
	t := biligo.NewCommClient(&biligo.CommSetting{})
	roomNum, err := strconv.ParseInt(room, 10, 64)
	if err != nil {
		return nil, err
	}
	info, err := t.LiveGetRoomInfoByID(roomNum)
	if err != nil {
		return nil, err
	}
	r, err := t.UserGetInfo(info.UID)
	if err != nil {
		return nil, err
	}
	return &model.RoomInfo{
		Status: r.LiveRoom.LiveStatus,
		Room:   strconv.FormatInt(info.RoomID, 10),
		Upper:  r.Name,
		Link:   r.LiveRoom.URL,
		Title:  r.LiveRoom.Title,
	}, nil
}

func (c *base) Stop() {

}
