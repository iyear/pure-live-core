package util

import "github.com/iyear/pure-live/conf"

var qn2desc = map[int]string{
	conf.QnBest: "原画",
	conf.QnHigh: "蓝光",
	conf.QnMid:  "超清",
	conf.QnLow:  "流畅",
}

func Qn2Desc(qn int) string {
	return qn2desc[qn]
}
