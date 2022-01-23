package abstract

import (
	"errors"
	"github.com/iyear/pure-live-core/model"
	"github.com/iyear/pure-live-core/pkg/conf"
)

type Client struct{}

// Plat return live platform id
func (c *Client) Plat() string {
	return conf.PlatAbstract
}

// GetPlayURL get live stream info
func (c *Client) GetPlayURL(room string, qn int) (*model.PlayURL, error) {
	_ = room
	_ = qn
	return nil, errors.New("please override the GetPlayURL function")
}

// GetRoomInfo get live room info
func (c *Client) GetRoomInfo(room string) (*model.RoomInfo, error) {
	_ = room
	return nil, errors.New("please override the GetRoomInfo function")
}

// Host return websocket host
func (c *Client) Host(room string) string {
	_ = room
	return ""
}

// Enter return data of entering the live room
func (c *Client) Enter(room string) (tp int, data [][]byte, err error) {
	_ = room
	return 0, nil, errors.New("please override the Enter function")
}

// Handle return websocket parsing data
func (c *Client) Handle(tp int, data []byte) (msg []model.Msg, matched bool, err error) {
	_ = tp
	_ = data
	return nil, false, errors.New("please override the Handle function")
}

// HeartBeat return heartbeat msg data
func (c *Client) HeartBeat() (tp int, data []byte, err error) {
	return 0, nil, errors.New("please override the HeartBeat function")
}

// SendDanmaku send danmaku
func (c *Client) SendDanmaku(room string, content string, tp int, color int64) error {
	_ = room
	_ = content
	_ = tp
	_ = color
	return errors.New("please override the SendDanmaku function")
}

// Stop stop and release resource
func (c *Client) Stop() {}
