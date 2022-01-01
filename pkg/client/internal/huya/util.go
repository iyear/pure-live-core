package huya

import (
	"fmt"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
	"github.com/iyear/pure-live-core/pkg/request"
	"github.com/tidwall/gjson"
	"regexp"
)

func getEnterMsg(lYyid, lChannelId, lSubChannelId int64) []byte {
	// TODO use tars to send enter msg
	buf := codec.NewBuffer()
	_ = buf.Write_int64(lYyid, 0)
	_ = buf.Write_bool(true, 1)
	_ = buf.Write_string("", 2)
	_ = buf.Write_string("", 3)
	_ = buf.Write_int64(lChannelId, 4)
	_ = buf.Write_int64(lSubChannelId, 5)
	_ = buf.Write_int64(0, 6)
	_ = buf.Write_int64(0, 7)

	// fmt.Println(hex.Dump(buf.ToBytes()))

	out := codec.NewBuffer()
	_ = out.Write_int32(1, 0)
	_ = out.WriteHead(13, 1)
	_ = out.WriteHead(0, 0)
	_ = out.Write_int32(int32(len(buf.ToBytes())), 0)

	return append(out.ToBytes(), buf.ToBytes()...)
}
func getRoomInfo(room string) (gjson.Result, error) {
	html := ""
	err := request.HTTP().GET(fmt.Sprintf("https://m.huya.com/%s", room)).SetHeader(H{
		"Content-Type": "application/x-www-form-urlencoded",
		"User-Agent":   "Mozilla/5.0 (Linux; Android 5.0; SM-G900P Build/LRX21T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Mobile Safari/537.36",
	}).BindBody(&html).Do()
	if err != nil {
		return gjson.Result{}, err
	}
	j := regexp.MustCompile(`{"roomProfile":(.*)}`).FindString(html)
	return gjson.Parse(j), nil
}
