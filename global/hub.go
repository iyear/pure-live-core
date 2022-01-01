package global

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/iyear/pure-live-core/model"
	"sync"
)

var Hub = &hub{}

type hub struct {
	Conn sync.Map
}
type Conn struct {
	Room   string          `json:"room"`
	Server *websocket.Conn `json:"ws"`
	Client model.Client    `json:"client"`
}

func GetConn(id string) (*Conn, error) {
	c, ok := Hub.Conn.Load(id)
	if !ok {
		return nil, fmt.Errorf("conn not found")
	}
	return c.(*Conn), nil
}
