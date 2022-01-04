package bilibili

import (
	"github.com/iyear/biligo"
	"github.com/iyear/pure-live-core/model"
	"github.com/iyear/pure-live-core/pkg/client/internal/abstract"
	"github.com/iyear/pure-live-core/pkg/conf"
	"github.com/iyear/pure-live-core/pkg/util"
	"strconv"
)

type base struct {
	*abstract.Client
}

// NewBiliBili .
func NewBiliBili() (model.Client, error) {
	if !conf.Account.BiliBili.Enable {
		return &BiliComm{
			client: biligo.NewCommClient(&biligo.CommSetting{
				DebugMode: false,
			}),
		}, nil
	}
	b, err := biligo.NewBiliClient(&biligo.BiliSetting{
		Auth: &biligo.CookieAuth{
			DedeUserID:      conf.Account.BiliBili.DedeUserID,
			DedeUserIDCkMd5: conf.Account.BiliBili.DedeUserIDCkMd5,
			SESSDATA:        conf.Account.BiliBili.SESSDATA,
			BiliJCT:         conf.Account.BiliBili.BiliJCT,
		},
		DebugMode: false,
	})
	if err != nil {
		return nil, err
	}
	return &BiliBili{client: b}, nil
}

// Plat .
func (c *base) Plat() string {
	return conf.PlatBiliBili
}

// GetPlayURL .
func (c *base) GetPlayURL(room string, qn int) (*model.PlayURL, error) {
	client := biligo.NewCommClient(&biligo.CommSetting{})
	roomNum, err := strconv.ParseInt(room, 10, 64)
	if err != nil {
		return nil, err
	}

	// 内部维护一个qn映射表
	q := map[int]int{
		conf.QnBest: 20000,
		conf.QnHigh: 10000,
		conf.QnMid:  400,
		conf.QnLow:  250,
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

// GetRoomInfo .
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

// Stop .
func (c *base) Stop() {

}
