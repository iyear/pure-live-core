package egame

import (
	"github.com/guonaihong/gout"
	"github.com/tidwall/gjson"
	"strings"
)

func getRoomInfo(room string) (gjson.Result, error) {
	resp := ""
	tmpl := `{"0":{"module":"pgg_live_read_svr","method":"get_live_and_profile_info","param":{"anchor_id":{{room}},"layout_id":"hot","index":1,"other_uid":0}}}`

	err := gout.GET("https://share.egame.qq.com/cgi-bin/pgg_async_fcgi").
		SetQuery(gout.H{
			"param": strings.ReplaceAll(tmpl, "{{room}}", room),
		}).BindBody(&resp).Do()

	if err != nil {
		return gjson.Result{}, err
	}

	return getRetData(resp), nil
}
func getWSToken(room string) (string, error) {
	info, err := getRoomInfo(room)
	if err != nil {
		return "", err
	}

	pid := info.Get("video_info.pid").String()

	resp := ""
	tmpl := `{"0":{"module":"pgg.ws_token_go_svr.DefObj","method":"get_token","param":{"scene_flag":16,"subinfo":{"page":{"scene":1,"page_id":{{room}},"str_id":"{{pid}}","msg_type_list":[1,2]}},"version":1,"message_seq":-1,"dc_param":{"params":{"info":{"aid":"{{room}}"}},"position":{"page_id":"QG_HEARTBEAT_PAGE_LIVE_ROOM"},"refer":{}},"other_uid":0}}}`

	err = gout.POST("https://share.egame.qq.com/cgi-bin/pgg_async_fcgi").
		SetWWWForm(gout.H{
			"param": strings.NewReplacer("{{room}}", room, "{{pid}}", pid).Replace(tmpl),
		}).BindBody(&resp).Do()
	if err != nil {
		return "", err
	}

	return getRetData(resp).Get("token").String(), nil
}
func getRetData(s string) gjson.Result {
	return gjson.Get(s, "data.\\0.retBody.data")
}
