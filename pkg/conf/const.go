package conf

import "errors"

const (
	PlatAbstract = "abstract"
	PlatBiliBili = "bilibili"
	PlatHuya     = "huya"
	PlatDouyu    = "douyu"
	PlatEGame    = "egame"
	PlatInke     = "inke"
)

const (
	EventDanmaku = "danmaku"
	EventCheck   = "check"
	EventHot     = "hot"
)

const (
	DanmakuTypeRight  = 0
	DanmakuTypeTop    = 1
	DanmakuTypeBottom = 2
)

const (
	QnBest = iota
	QnHigh
	QnMid
	QnLow
)

const (
	StreamFlv  = "flv"
	StreamHls  = "hls"
	StreamM3U8 = "m3u8"
)

// biz key

const (
	BizRoomInfo = "ri"
)

var (
	ErrSkip = errors.New("skip")
)
