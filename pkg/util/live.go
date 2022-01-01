package util

import "github.com/iyear/pure-live-core/pkg/conf"

var qn2desc = map[int]string{
	conf.QnBest: "原画",
	conf.QnHigh: "蓝光",
	conf.QnMid:  "超清",
	conf.QnLow:  "流畅",
}

// Qn2Desc returns the description of quality.
func Qn2Desc(qn int) string {
	return qn2desc[qn]
}

var mode2desc = map[int]string{
	conf.DanmakuTypeTop:    "顶部",
	conf.DanmakuTypeRight:  "飞行",
	conf.DanmakuTypeBottom: "底部",
}

// DmMode2Desc returns the description of danmaku mode.
func DmMode2Desc(mode int) string {
	return mode2desc[mode]
}

var plat2desc = map[string]string{
	conf.PlatBiliBili: "哔哩哔哩",
	conf.PlatHuya:     "虎牙",
	conf.PlatDouyu:    "斗鱼",
	conf.PlatEGame:    "企鹅电竞",
}

func Plat2Desc(plat string) string {
	return plat2desc[plat]
}
