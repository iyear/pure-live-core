package egame

import (
	"fmt"
	"github.com/guonaihong/gout"
	"github.com/iyear/pure-live/conf"
	"github.com/iyear/pure-live/model"
	"github.com/iyear/pure-live/util"
	"github.com/tidwall/gjson"
	"strings"
)

type EGame struct{}

func NewEGame() (model.Client, error) {
	return &EGame{}, nil
}

func (e *EGame) Plat() string {
	return conf.PlatEGame
}

func getInfo(room string) (*gjson.Result, error) {
	resp := ""
	tmpl := `{"0":{"module":"pgg_live_read_svr","method":"get_live_and_profile_info","param":{"anchor_id":{{id}},"layout_id":"hot","index":1,"other_uid":0}}}`

	err := gout.GET("https://share.egame.qq.com/cgi-bin/pgg_async_fcgi").SetQuery(gout.H{
		"param": strings.ReplaceAll(tmpl, "{{id}}", room),
	}).BindBody(&resp).Do()

	if err != nil {
		return nil, err
	}

	r := gjson.Get(resp, "data.\\0.retBody.data")
	return &r, nil

}
func (e *EGame) GetPlayURL(room string, qn int) (*model.PlayURL, error) {
	r, err := getInfo(room)
	if err != nil {
		return nil, err
	}
	url := r.Get("video_info.stream_infos.0.play_url").String()

	return &model.PlayURL{
		Qn:     conf.QnBest,
		Desc:   util.Qn2Desc(conf.QnBest),
		Origin: url,
		CORS:   true,
		Type:   conf.StreamFlv,
	}, nil
}

func (e *EGame) GetRoomInfo(room string) (*model.RoomInfo, error) {
	r, err := getInfo(room)
	if err != nil {
		return nil, err
	}
	title := r.Get("video_info.title").String()
	profile := r.Get("profile_info")
	return &model.RoomInfo{
		Status: int(profile.Get("is_live").Int()),
		Room:   room,
		Upper:  profile.Get("nick_name").String(),
		Link:   fmt.Sprintf("https://egame.qq.com/%s", room),
		Title:  title,
	}, nil
}

func (e *EGame) Host() string {
	panic("implement me")
}

func (e *EGame) Enter(room string) (tp int, data [][]byte, err error) {
	panic("implement me")
}

func (e *EGame) Handle(tp int, data []byte) (msg []model.Msg, matched bool, err error) {
	panic("implement me")
}

func (e *EGame) HeartBeat() (tp int, data []byte, err error) {
	panic("implement me")
}

func (e *EGame) SendDanmaku(room string, content string, tp int, color int64) error {
	panic("implement me")
}

func (e *EGame) Stop() {

}
