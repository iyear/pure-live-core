package egame

import (
	"encoding/hex"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/iyear/pure-live-core/model"
	"github.com/iyear/pure-live-core/pkg/client/internal/abstract"
	"github.com/iyear/pure-live-core/pkg/client/internal/egame/internal/packet"
	"github.com/iyear/pure-live-core/pkg/conf"
	"github.com/iyear/pure-live-core/pkg/util"
	"github.com/tidwall/gjson"
)

const hb = "000000120012000100070000000100000000"

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
func (e *EGame) Host(room string) string {
	_ = room
	return "wss://barragepush.egame.qq.com/sub"
}

// Enter .
func (e *EGame) Enter(room string) (int, [][]byte, error) {
	token, err := getWSToken(room)
	if err != nil {
		return -1, nil, err
	}

	body := util.PutBytes(
		[]byte{uint8(7)},
		util.BigEndianUint32(uint32(len(token))),
		[]byte(token),
	)

	header := util.PutBytes(
		util.BigEndianUint32(uint32(18+len(body))),
		util.BigEndianUint16(18),
		util.BigEndianUint16(1),
		util.BigEndianUint16(1),
		util.BigEndianUint32(0),
		util.BigEndianUint16(0),
		util.BigEndianUint16(0),
	)

	data := append(header, body...)
	return websocket.BinaryMessage, [][]byte{data}, nil
}

// HeartBeat .
func (e *EGame) HeartBeat() (tp int, data []byte, err error) {
	b, err := hex.DecodeString(hb)
	if err != nil {
		return 0, nil, err
	}
	return websocket.BinaryMessage, b, nil
}

// Handle .
func (e *EGame) Handle(tp int, data []byte) ([]model.Msg, bool, error) {
	if tp != websocket.BinaryMessage {
		return nil, false, nil
	}
	resp, err := packet.Decode(data)
	if err != nil {
		return nil, false, err
	}

	// t, _ := json.Marshal(resp)
	// fmt.Println(string(t))

	if resp.Operation != 3 {
		return nil, false, nil
	}

	var msgs []model.Msg
	for _, body := range resp.Body {
		switch body.MsgType {
		case 1:
			for _, bd := range body.BinData {
				result := gjson.Parse(string(bd))
				if result.Get("type").Int() == 0 {
					msgs = append(msgs, &model.MsgDanmaku{
						Content: result.Get("content").String(),
						Type:    conf.DanmakuTypeRight,
						Color:   16777215,
					})
				}
			}
		}
	}
	return msgs, true, nil
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
