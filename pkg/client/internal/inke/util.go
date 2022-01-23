package inke

import (
	"errors"
	"fmt"
	"github.com/iyear/pure-live-core/pkg/request"
	"github.com/tidwall/gjson"
)

func getRoomInfo(room string) (gjson.Result, error) {
	body := ""
	err := request.HTTP().
		GET(fmt.Sprintf("https://webapi.busi.inke.cn/web/live_share_pc?uid=%s", room)).
		BindBody(&body).
		Do()
	if err != nil {
		return gjson.Result{}, err
	}

	info := gjson.Parse(body)
	if info.Get("error_code").Int() != 0 {
		return gjson.Result{}, errors.New(info.Get("message").String())
	}

	return info, nil
}

func getRoomLink(room string) string {
	return fmt.Sprintf("https://www.inke.cn/liveroom/index.html?uid=%s", room)
}
