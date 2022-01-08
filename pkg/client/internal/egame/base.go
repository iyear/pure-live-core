package egame

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/iyear/pure-live-core/model"
	"github.com/iyear/pure-live-core/pkg/client/internal/abstract"
	"github.com/iyear/pure-live-core/pkg/conf"
	"github.com/iyear/pure-live-core/pkg/util"
)

type EGame struct {
	*abstract.Client
}

// NewEGame .
func NewEGame() (model.Client, error) {
	return &EGame{}, nil
}

// Plat .
func (e *EGame) Plat() string {
	return conf.PlatEGame
}

// GetPlayURL .
func (e *EGame) GetPlayURL(room string, qn int) (*model.PlayURL, error) {
	r, err := getRoomInfo(room)
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

// GetRoomInfo .
func (e *EGame) GetRoomInfo(room string) (*model.RoomInfo, error) {
	r, err := getRoomInfo(room)
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

// Host .
func (e *EGame) Host() string {
	return "wss://barragepush.egame.qq.com/sub"
}

// Enter .
func (e *EGame) Enter(room string) (int, [][]byte, error) {
	token, err := getWSToken(room)
	if err != nil {
		return -1, nil, err
	}

	body := append([]byte{uint8(7)}, util.BigEndianUint32(uint32(len(token)))...)
	body = append(body, []byte(token)...)

	header := append(util.BigEndianUint32(uint32(18+len(body))), util.BigEndianUint16(18)...)
	header = append(header, util.BigEndianUint16(1)...)
	header = append(header, util.BigEndianUint16(1)...)
	header = append(header, util.BigEndianUint32(0)...)
	header = append(header, util.BigEndianUint16(0)...)
	header = append(header, util.BigEndianUint16(0)...)

	data := append(header, body...)
	return websocket.BinaryMessage, [][]byte{data}, nil
}

// Handle .
func (e *EGame) Handle(tp int, data []byte) (msg []model.Msg, matched bool, err error) {
	_ = tp
	_ = data
	return nil, false, nil
}

// HeartBeat .
func (e *EGame) HeartBeat() (tp int, data []byte, err error) {
	return 0, nil, fmt.Errorf("not supported")
}

// SendDanmaku .
func (e *EGame) SendDanmaku(room string, content string, tp int, color int64) error {
	_ = room
	_ = content
	_ = tp
	_ = color
	return fmt.Errorf("not supported")
}

// Stop .
func (e *EGame) Stop() {

}
