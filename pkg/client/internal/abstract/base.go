package abstract

import (
	"errors"
	"github.com/iyear/pure-live/model"
	"github.com/iyear/pure-live/pkg/conf"
)

type Client struct{}

func (c *Client) Plat() string {
	return conf.PlatAbstract
}

func (c *Client) GetPlayURL(room string, qn int) (*model.PlayURL, error) {
	_ = room
	_ = qn
	return nil, errors.New("please override the GetPlayURL function")
}

func (c *Client) GetRoomInfo(room string) (*model.RoomInfo, error) {
	_ = room
	return nil, errors.New("please override the GetRoomInfo function")
}

func (c *Client) Host() string {
	return ""
}

func (c *Client) Enter(room string) (tp int, data [][]byte, err error) {
	_ = room
	return 0, nil, errors.New("please override the Enter function")
}

func (c *Client) Handle(tp int, data []byte) (msg []model.Msg, matched bool, err error) {
	_ = tp
	_ = data
	return nil, false, errors.New("please override the Handle function")
}

func (c *Client) HeartBeat() (tp int, data []byte, err error) {
	return 0, nil, errors.New("please override the HeartBeat function")
}

func (c *Client) SendDanmaku(room string, content string, tp int, color int64) error {
	_ = room
	_ = content
	_ = tp
	_ = color
	return errors.New("please override the SendDanmaku function")
}

func (c *Client) Stop() {}
