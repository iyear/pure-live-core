## 如何新增一个Client?

首先你需要实现以下接口:
```go
type Client interface {
	Plat() string
	// GetPlayURL qn传入 conf.QnBest conf.QnHigh conf.QnMid conf.QnLow
	GetPlayURL(room string, qn int) (*PlayURL, error)
	// GetRoomInfo room可以为短号也可以为长号
	GetRoomInfo(room string) (*RoomInfo, error)
	// Host ws host
	Host() string
	// Enter 一次可以返回多条消息，hub将按顺序依次发送，用于需要一次发送多条进入直播间消息的场景
	Enter(room string) (tp int, data [][]byte, err error)
	// Handle matched为是否读取msg的操作，用于跳过不想匹配的消息，err为错误，先判断错误再判断matched
	Handle(tp int, data []byte) (msg []Msg, matched bool, err error)
	HeartBeat() (tp int, data []byte, err error)
	// SendDanmaku tp 1:top 0:right 2:bottom color:十进制颜色值
	SendDanmaku(room string, content string, tp int, color int64) error
	// Stop 释放内部资源
	Stop()
}
```

- 所有的 `tp` 为 `Websocket Message Type`，使用 `websocket.xxxMessage` 传入

具体参考已有的Client写法

1. 新增后需要在 `/conf/const.go` 新增平台名常量
2. 在 `/docs/API.md` 中新增平台的说明
3. 需要在 `/pkg/client/plat.go` 中新增一个 `Client` 的实现
