package egame

import (
	"fmt"
	"github.com/guonaihong/gout"
	"github.com/iyear/pure-live-core/model"
	"github.com/iyear/pure-live-core/pkg/client/internal/abstract"
	"github.com/iyear/pure-live-core/pkg/conf"
	"github.com/iyear/pure-live-core/pkg/util"
	"github.com/tidwall/gjson"
	"strings"
)

type EGame struct {
	*abstract.Client
}

// NewEGame
func NewEGame() (model.Client, error) {
	return &EGame{}, nil
}

// Plat
func (e *EGame) Plat() string {
	return conf.PlatEGame
}

func getInfo(room string) (*gjson.Result, error) {
	resp := ""
	tmpl := `{"0":{"module":"pgg_live_read_svr","method":"get_live_and_profile_info","param":{"anchor_id":{{id}},"layout_id":"hot","index":1,"other_uid":0}}}`

	err := gout.GET("https://share.egame.qq.com/cgi-bin/pgg_async_fcgi").
		SetQuery(gout.H{
			"param": strings.ReplaceAll(tmpl, "{{id}}", room),
		}).BindBody(&resp).Do()

	if err != nil {
		return nil, err
	}

	r := gjson.Get(resp, "data.\\0.retBody.data")
	return &r, nil

}

// GetPlayURL
func (e *EGame) GetPlayURL(room string, qn int) (*model.PlayURL, error) {
	r, err := getInfo(room)
	if err != nil {
		return nil, err
	}
	url := r.Get("video_info.stream_infos.0.play_url").String()

	return &model.PlayURL{
		Qn:     qn,
		Desc:   util.Qn2Desc(qn),
		Origin: url,
		CORS:   true,
		Type:   conf.StreamFlv,
	}, nil
}

// GetRoomInfo
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

// Host
func (e *EGame) Host() string {
	return "wss://barragepush.egame.qq.com/sub"
}

// Enter
func (e *EGame) Enter(room string) (tp int, data [][]byte, err error) {
	_ = room
	return 0, nil, fmt.Errorf("not supported")
}

// Handle
func (e *EGame) Handle(tp int, data []byte) (msg []model.Msg, matched bool, err error) {
	_ = tp
	_ = data
	return nil, false, nil
}

// HeartBeat
func (e *EGame) HeartBeat() (tp int, data []byte, err error) {
	return 0, nil, fmt.Errorf("not supported")
}

// SendDanmaku
func (e *EGame) SendDanmaku(room string, content string, tp int, color int64) error {
	_ = room
	_ = content
	_ = tp
	_ = color
	return fmt.Errorf("not supported")
}

// Stop
func (e *EGame) Stop() {

}
