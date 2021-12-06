package request

import (
	"github.com/guonaihong/gout"
	"github.com/guonaihong/gout/dataflow"
	"github.com/iyear/pure-live/conf"
	"github.com/iyear/pure-live/util"
	"net/http"
)

func HTTP() *dataflow.DataFlow {
	c := http.DefaultClient
	tsp := &http.Transport{}
	if conf.C.Socks5.Enable {
		tsp.Dial = util.MustGetSocks5(conf.C.Socks5.Host, conf.C.Socks5.Port, conf.C.Socks5.User, conf.C.Socks5.Password).Dial
	}
	c.Transport = tsp
	return gout.New(c).Debug(false)
}
