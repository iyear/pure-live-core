package huya

import (
	"fmt"
	"github.com/iyear/pure-live-core/pkg/request"
	"github.com/tidwall/gjson"
	"regexp"
)

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
