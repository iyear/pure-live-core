package egame

import (
	"fmt"
	"github.com/guonaihong/gout"
	"github.com/tidwall/gjson"
	"log"
	"strings"
	"testing"
)

// room_url = 'https://share.egame.qq.com/cgi-bin/pgg_async_fcgi'
//        post_data = {
//            'param': '''{"0":{"module":"pgg_live_read_svr","method":"get_live_and_profile_info","param":{"anchor_id":'''
//                     + str(self.rid) + ''',"layout_id":"hot","index":1,"other_uid":0}}}'''
//        }
//        try:
//            response = requests.post(url=room_url, data=post_data).json()
//            data = response.get('data', 0)
//            if data:
//                video_info = data.get('0').get(
//                    'retBody').get('data').get('video_info')
//                pid = video_info.get('pid', 0)
//                if pid:
//                    is_live = data.get('0').get(
//                    'retBody').get('data').get('profile_info').get('is_live', 0)
//                    if is_live:
//                        play_url = video_info.get('stream_infos')[
//                            0].get('play_url')
//                        real_url = re.findall(r'([\w\W]+?)&uid=', play_url)[0]
//                    else:
//                        raise Exception('直播间未开播')
//                else:
//                    raise Exception('直播间未启用')
//            else:
//                raise Exception('直播间不存在')
//        except:
//            raise Exception('数据请求错误')
//        return real_url
func TestEGame(t *testing.T) {
	resp := ""
	tmpl := `{"0":{"module":"pgg_live_read_svr","method":"get_live_and_profile_info","param":{"anchor_id":[[id]],"layout_id":"hot","index":1,"other_uid":0}}}`
	err := gout.GET("https://share.egame.qq.com/cgi-bin/pgg_async_fcgi").SetQuery(gout.H{
		"param": strings.ReplaceAll(tmpl, "[[id]]", "436309399"),
	}).BindBody(&resp).Do()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(resp)
	// http://3954-out.liveplay.myqcloud.com/live/3954_383204988_1024p.flv?bizid=3954&txSecret=594c5267d2d345ee61aaa7084ed73757&txTime=61c55b8e&uid=0&fromdj=&_qedj_t=BhgX1Z1373Muj17v4RWDFG7xFNsZviDJw7EQASYUMzgzMjA0OTg4XzE2Mzk3MTkxODEyYbw530xWHHBnSzVDX1FMQUFDajd3MWFVdE1GQU1EMkxRTmhmAHyGK3BnZ19saXZlX3JlYWRfaWZjX210X3N2ci5lbnRyeV9oNV9saXZlX3Jvb22WAKYNNjAuMTkxLjEyMi4zNLzGANYA7PwP
	r := gjson.Get(resp, "data.\\0.retBody.data")

	video := r.Get("video_info")
	pid := video.Get("pid").String()
	if pid == "" {
		log.Println("直播间不存在")
		return
	}

	url := video.Get("stream_infos.0.play_url")

	profile := r.Get("profile_info")

	status := int(profile.Get("is_live").Int())

	upper := profile.Get("nick_name").String()

	fmt.Println(status, upper)
	fmt.Println(url)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestEGameWs(t *testing.T) {
	room := "436309399"
	r, err := getRoomInfo(room)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	pid := r.Get("video_info.pid").String()

	// params = {
	//         'param': json.dumps({"0":{"module":"pgg.ws_token_go_svr.DefObj","method":"get_token","param":{"scene_flag":16,"subinfo":{"page":{"scene":1,"page_id":int(page_id),"str_id":str(str_id),"msg_type_list":[1,2]}},"version":1,"message_seq":-1,"dc_param":{"params":{"info":{"aid":aid}},"position":{"page_id":"QG_HEARTBEAT_PAGE_LIVE_ROOM"},"refer":{}},"other_uid":0}}})
	//                }
	html := ""
	tmpl := `{"0":{"module":"pgg.ws_token_go_svr.DefObj","method":"get_token","param":{"scene_flag":16,"subinfo":{"page":{"scene":1,"page_id":{{room}},"str_id":"{{pid}}","msg_type_list":[1,2]}},"version":1,"message_seq":-1,"dc_param":{"params":{"info":{"aid":"{{room}}"}},"position":{"page_id":"QG_HEARTBEAT_PAGE_LIVE_ROOM"},"refer":{}},"other_uid":0}}}`
	param := strings.NewReplacer("{{room}}", room, "{{pid}}", pid).Replace(tmpl)

	fmt.Println(param)

	_ = gout.POST("https://share.egame.qq.com/cgi-bin/pgg_async_fcgi").
		SetWWWForm(gout.H{
			"param": param,
		}).BindBody(&html).Do()

	fmt.Println(html)
}
