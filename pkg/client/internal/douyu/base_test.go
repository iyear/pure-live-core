package douyu

import (
	"crypto/md5"
	"fmt"
	"github.com/dop251/goja"
	"github.com/guonaihong/gout"
	"github.com/iyear/pure-live-core/pkg/conf"
	"github.com/iyear/pure-live-core/pkg/request"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestDouyu_GetPlayURL3(t *testing.T) {
	douyu := Douyu{}
	url, err := douyu.GetPlayURL("475252", conf.QnBest)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fmt.Println(url)
}
func TestDouyu_GetRoomInfo(t *testing.T) {
	douyu := Douyu{}
	info, err := douyu.GetRoomInfo("475252")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fmt.Println(info)
}
func TestDouyu_GetPlayURL(t *testing.T) {
	rid := "1126960"
	did := "10000000000000000000000000001501"
	t10 := strconv.FormatInt(time.Now().Unix(), 10)
	html := ""
	err := request.HTTP().GET(fmt.Sprintf("https://m.douyu.com/%s", rid)).BindBody(&html).Do()
	if err != nil {
		return
	}
	fmt.Println(html)
	jsUb9 := regexp.MustCompile(`(function ub98484234.*)\s(var.*)`).FindString(html)
	jsUb9 = regexp.MustCompile("eval.*;}").ReplaceAllString(jsUb9, `strc;}`)

	vm := goja.New()

	if _, err = vm.RunString(jsUb9); err != nil {
		log.Println(err)
		return
	}
	ub9, ok := goja.AssertFunction(vm.Get("ub98484234"))
	if !ok {
		return
	}
	res, err := ub9(goja.Undefined())
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(res)

	value := regexp.MustCompile(`v=(\d+)`).FindAllStringSubmatch(res.String(), -1)[0][1]
	fmt.Println(value)

	// rb = DouYu.md5(self.rid + self.did + self.t10 + v)
	rb := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%s%s%s", rid, did, t10, value))))
	fmt.Println(rb)
	// func_sign = re.sub(r'return rt;}\);?', 'return rt;}', res)
	//        func_sign = func_sign.replace('(function (', 'function sign(')
	//        func_sign = func_sign.replace('CryptoJS.MD5(cb).toString()', '"' + rb + '"')
	funcSign := regexp.MustCompile(`return rt;}\);?`).ReplaceAllString(res.String(), `return rt;}`)
	funcSign = strings.ReplaceAll(funcSign, `(function (`, `function sign(`)
	funcSign = strings.ReplaceAll(funcSign, `CryptoJS.MD5(cb).toString()`, `"`+rb+`"`)

	fmt.Println(funcSign)
	// js = execjs.compile(func_sign)
	//        params = js.call('sign', self.rid, self.did, self.t10)
	//        params += '&ver=219032101&rid={}&rate=-1'.format(self.rid)
	//
	//        url = 'https://m.douyu.com/api/room/ratestream'
	//        res = self.s.post(url, params=params).text
	//        key = re.search(r'(\d{1,8}[0-9a-zA-Z]+)_?\d{0,4}(.m3u8|/playlist)', res).group(1)

	if _, err = vm.RunString(funcSign); err != nil {
		log.Println(err)
		return
	}
	sign, ok := goja.AssertFunction(vm.Get("sign"))
	if !ok {
		log.Println(false)
		return
	}
	param, err := sign(goja.Undefined(), vm.ToValue(rid), vm.ToValue(did), vm.ToValue(t10))
	if err != nil {
		log.Println(err)
		return
	}
	// params := fmt.Sprintf("%s&ver=219032101&rid=%s&rate=0", param.String(), rid)
	params := fmt.Sprintf("%s&cdn=ws-h5&rate=0", param)
	fmt.Println(params)
	resp := ""
	err = request.HTTP().POST("https://m.douyu.com/api/room/ratestream?" + params).
		SetHeader(gout.H{
			"UserAgent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36",
			"referer":   "https://www.douyu.com/",
			"origin":    "https://www.douyu.com",
		}).
		BindBody(&resp).
		Do()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(resp)
}
func TestDouyu_GetPlayURL2(t *testing.T) {
	log.SetFlags(log.Lshortfile)
	rid := "9408500"
	did := "10000000000000000000000000001501"
	t10 := strconv.FormatInt(time.Now().Unix(), 10)
	html := ""
	err := gout.GET(fmt.Sprintf("https://www.douyu.com/%s", rid)).BindBody(&html).Do()
	if err != nil {
		return
	}
	fmt.Println(html)
	jsUb9 := regexp.MustCompile(`(vdwdae325w_64we[\s\S]*function ub98484234[\s\S]*?)function`).FindString(html)
	log.Println(jsUb9)
	jsUb9 = strings.TrimSuffix(jsUb9, "function")
	log.Println(jsUb9)
	jsUb9 = regexp.MustCompile(`eval.*?;}`).ReplaceAllString(jsUb9, `strc;}`)
	log.Println(jsUb9)
	vm := goja.New()

	if _, err = vm.RunString(jsUb9); err != nil {
		log.Println(err)
		return
	}
	ub9, ok := goja.AssertFunction(vm.Get("ub98484234"))
	if !ok {
		log.Println(9999)
		return
	}
	res, err := ub9(goja.Undefined())
	if err != nil {
		log.Println(err)
		return
	}

	// fmt.Println(res)

	value := regexp.MustCompile(`v=(\d+)`).FindAllStringSubmatch(res.String(), -1)[0][1]
	// fmt.Println(value)

	rb := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%s%s%s", rid, did, t10, value))))
	// fmt.Println(rb)
	funcSign := regexp.MustCompile(`return rt;}\);?`).ReplaceAllString(res.String(), `return rt;}`)
	funcSign = strings.ReplaceAll(funcSign, `(function (`, `function sign(`)
	funcSign = strings.ReplaceAll(funcSign, `CryptoJS.MD5(cb).toString()`, `"`+rb+`"`)

	// fmt.Println(funcSign)
	if _, err = vm.RunString(funcSign); err != nil {
		log.Println(err)
		return
	}
	sign, ok := goja.AssertFunction(vm.Get("sign"))
	if !ok {
		log.Println(false)
		return
	}
	param, err := sign(goja.Undefined(), vm.ToValue(rid), vm.ToValue(did), vm.ToValue(t10))
	if err != nil {
		log.Println(err)
		return
	}
	// params := fmt.Sprintf("%s&ver=219032101&rid=%s&rate=0", param.String(), rid)
	params := fmt.Sprintf("%s&cdn=ws-h5&rate=0", param)
	// fmt.Println(params)
	resp := ""
	err = gout.POST(fmt.Sprintf("https://www.douyu.com/lapi/live/getH5Play/%s?", rid) + params).
		SetHeader(gout.H{
			"UserAgent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36",
			"referer":   "https://www.douyu.com/",
			"origin":    "https://www.douyu.com",
		}).
		BindBody(&resp).
		Do()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(resp)
}
func TestGet(t *testing.T) {

	flv, err := os.Create("./1.flv")
	if err != nil {
		log.Println(err)
		return
	}
	resp, err := http.Get("http://dyscdnali1.douyucdn.cn/live/606118r6n9kQ9krJ.flv?uuid=")
	if err != nil {
		log.Println(err)
		return
	}
	ti := time.NewTimer(10 * time.Second)
	defer ti.Stop()
	var count int64 = 0
	var buf = make([]byte, 4096)
	for {
		select {
		case <-ti.C:
			return
		default:
			n, err := resp.Body.Read(buf)
			if err != nil || n == 0 {
				fmt.Println("出现错误", err)
				return
			}

			if _, err = flv.WriteAt(buf, count); err != nil {
				log.Println(err)
				return
			}
			count += int64(n)
		}
	}
}
