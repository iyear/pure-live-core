package util

import "github.com/iyear/pure-live/pkg/conf"

var qn2desc = map[int]string{
	conf.QnBest: "原画",
	conf.QnHigh: "蓝光",
	conf.QnMid:  "超清",
	conf.QnLow:  "流畅",
}

func Qn2Desc(qn int) string {
	return qn2desc[qn]
}

var mode2desc = map[int]string{
	conf.DanmakuTypeTop:    "顶部",
	conf.DanmakuTypeRight:  "飞行",
	conf.DanmakuTypeBottom: "底部",
}

func DmMode2Desc(mode int) string {
	return mode2desc[mode]
}
