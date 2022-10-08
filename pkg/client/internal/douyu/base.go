package douyu

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/dop251/goja"
	"github.com/gorilla/websocket"
	"github.com/guonaihong/gout"
	"github.com/iyear/pure-live-core/model"
	"github.com/iyear/pure-live-core/pkg/client/internal/abstract"
	"github.com/iyear/pure-live-core/pkg/conf"
	"github.com/iyear/pure-live-core/pkg/request"
	"github.com/iyear/pure-live-core/pkg/util"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Douyu struct {
	*abstract.Client
}

func NewDouyu() (model.Client, error) {
	return &Douyu{}, nil
}

// Plat
func (d *Douyu) Plat() string {
	return conf.PlatDouyu
}

// GetPlayURL
// cdn: 主线路ws-h5、备用线路tct-h5pot
// rate: 1流畅；2高清；3超清；4蓝光4M；0蓝光8M或10M
func (d *Douyu) GetPlayURL(room string, qn int) (*model.PlayURL, error) {
	var qnm = map[int]int{
		conf.QnBest: 0,
		conf.QnHigh: 4,
		conf.QnMid:  3,
		conf.QnLow:  1,
	}

	did := "10000000000000000000000000001501"
	t10 := strconv.FormatInt(time.Now().Unix(), 10)
	html := ""
	if err := request.HTTP().GET(fmt.Sprintf("https://www.douyu.com/%s", room)).BindBody(&html).Do(); err != nil {
		return nil, err
	}
	// fmt.Println(html)
	jsUb9 := regexp.MustCompile(`(vdwdae325w_64we[\s\S]*function ub98484234[\s\S]*?)function`).FindString(html)
	// fmt.Println(jsUb9)
	// 最后还有个function要去掉，不然编译不了
	jsUb9 = strings.TrimSuffix(jsUb9, "function")
	// fmt.Println(jsUb9)
	jsUb9 = regexp.MustCompile(`eval.*?;}`).ReplaceAllString(jsUb9, `strc;}`)
	// fmt.Println(jsUb9)
	vm := goja.New()

	if _, err := vm.RunString(jsUb9); err != nil {
		return nil, err
	}
	ub9, ok := goja.AssertFunction(vm.Get("ub98484234"))
	if !ok {
		return nil, fmt.Errorf("failed to assert function ub9")
	}
	res, err := ub9(goja.Undefined())
	if err != nil {
		return nil, err
	}

	// fmt.Println(res)

	value := regexp.MustCompile(`v=(\d+)`).FindAllStringSubmatch(res.String(), -1)[0][1]
	// fmt.Println(value)

	rb := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%s%s%s", room, did, t10, value))))
	// fmt.Println(rb)
	funcSign := regexp.MustCompile(`return rt;}\);?`).ReplaceAllString(res.String(), `return rt;}`)
	funcSign = strings.ReplaceAll(funcSign, `(function (`, `function sign(`)
	funcSign = strings.ReplaceAll(funcSign, `CryptoJS.MD5(cb).toString()`, `"`+rb+`"`)

	// fmt.Println(funcSign)
	if _, err = vm.RunString(funcSign); err != nil {
		return nil, err
	}
	sign, ok := goja.AssertFunction(vm.Get("sign"))
	if !ok {
		return nil, fmt.Errorf("failed to assert function sign")
	}
	param, err := sign(goja.Undefined(), vm.ToValue(room), vm.ToValue(did), vm.ToValue(t10))
	if err != nil {
		return nil, err
	}
	// params := fmt.Sprintf("%s&ver=219032101&rid=%s&rate=0", param.String(), rid)
	params := fmt.Sprintf("%s&cdn=ws-h5&rate=%d", param, qnm[qn])
	// fmt.Println(params)
	var resp struct {
		Error int    `json:"error"`
		Msg   string `json:"msg"`
		Data  struct {
			RoomID       int64  `json:"room_id"`
			IsMixed      bool   `json:"is_mixed"`
			MixedLive    string `json:"mixed_live"`
			MixedURL     string `json:"mixed_url"`
			RtmpCdn      string `json:"rtmp_cdn"`
			RtmpURL      string `json:"rtmp_url"`
			RtmpLive     string `json:"rtmp_live"`
			ClientIP     string `json:"client_ip"`
			InNA         int    `json:"inNA"`
			RateSwitch   int    `json:"rateSwitch"`
			Rate         int    `json:"rate"`
			CdnsWithName []*struct {
				Name   string `json:"name"`
				Cdn    string `json:"cdn"`
				IsH265 bool   `json:"isH265"`
			} `json:"cdnsWithName"`
			Multirates []*struct {
				Name    string `json:"name"`
				Rate    int    `json:"rate"`
				HighBit int    `json:"highBit"`
				Bit     int    `json:"bit"`
			} `json:"multirates"`
		}
	}

	err = request.HTTP().POST(fmt.Sprintf("https://www.douyu.com/lapi/live/getH5Play/%s?", room) + params).
		SetHeader(gout.H{
			"UserAgent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36",
			"referer":   "https://www.douyu.com/",
			"origin":    "https://www.douyu.com",
		}).
		BindJSON(&resp).
		Do()
	if err != nil {
		return nil, err
	}
	// fmt.Println(resp)
	if resp.Error != 0 {
		return nil, fmt.Errorf("failed to get play url: %s", resp.Msg)
	}
	return &model.PlayURL{
		Qn:     qn,
		Desc:   util.Qn2Desc(qn),
		Origin: fmt.Sprintf("http://akm-tct.douyucdn.cn/live/%s?uuid=", strings.Split(resp.Data.RtmpLive, "?")[0]),
		CORS:   true,
		Type:   conf.StreamFlv,
	}, nil
}

// GetRoomInfo 通过房间号获取房间信息
func (d *Douyu) GetRoomInfo(room string) (*model.RoomInfo, error) {
	var info struct {
		Error int `json:"error"`
		Data  struct {
			RoomId     string `json:"room_id"`
			OwnerName  string `json:"owner_name"`
			RoomStatus string `json:"room_status"`
			RoomName   string `json:"room_name"`
		} `json:"data"`
	}
	if err := request.HTTP().GET(fmt.Sprintf("https://open.douyucdn.cn/api/RoomApi/room/%s", room)).BindJSON(&info).Do(); err != nil {
		zap.S().Warnf("Douyu: GetRoomInfo: http.Get room:%v, err:%v", room, err)
		return nil, err
	}
	if info.Error != 0 {
		zap.S().Warnf("Douyu: GetRoomInfo: rsp err code not 0, room:%v", room)
		return nil, errors.New("request err")
	}
	link := fmt.Sprintf("https://www.douyu.com/%s", info.Data.RoomId)
	return &model.RoomInfo{
		Status: util.IF(info.Data.RoomStatus == "1", 1, 0).(int),
		Room:   info.Data.RoomId,
		Upper:  info.Data.OwnerName,
		Link:   link,
		Title:  info.Data.RoomName,
	}, nil
}

// Host
func (d *Douyu) Host(room string) string {
	_ = room
	return "wss://danmuproxy.douyu.com:8503/"
}

func (d *Douyu) Enter(room string) (int, [][]byte, error) {
	return websocket.BinaryMessage, [][]byte{
		encode(map[string]string{"type": "loginreq", "roomid": room}),
		encode(map[string]string{"type": "joingroup", "rid": room, "gid": "-9999"}),
	}, nil
}

func (d *Douyu) Handle(tp int, data []byte) ([]model.Msg, bool, error) {
	if tp != websocket.BinaryMessage || len(data) < 13 {
		return nil, false, nil
	}
	packs := decode(data)
	// fmt.Println(packs)
	var msgs []model.Msg
	for _, pack := range packs {
		tp := gjson.Get(pack, "type").String()
		switch tp {
		case "chatmsg":
			msgs = append(msgs, &model.MsgDanmaku{
				Content: gjson.Get(pack, "txt").String(),
				Type:    conf.DanmakuTypeRight, // 没找到弹幕显示位置的字段
				Color:   colorConv(gjson.Get(pack, "col").String()),
			})
		}

	}
	return msgs, true, nil
}

func (d *Douyu) HeartBeat() (int, []byte, error) {
	return websocket.BinaryMessage, encode(map[string]string{
		"type": "mrkl",
	}), nil
}

func (d *Douyu) SendDanmaku(room string, content string, tp int, color int64) error {
	_ = room
	_ = content
	_ = tp
	_ = color
	return fmt.Errorf("todo")
}

func (d *Douyu) Stop() {

}
