package request

import (
	"github.com/guonaihong/gout"
	"github.com/guonaihong/gout/dataflow"
	"github.com/iyear/pure-live-core/pkg/util"
	"net"
	"net/http"
)

var dial = net.Dial

// SetSocks5 set socks5 proxy
func SetSocks5(host string, port int, user, password string) {
	dial = util.MustGetSocks5(host, port, user, password).Dial
}

// HTTP http request
func HTTP() *dataflow.DataFlow {
	c := http.DefaultClient
	tsp := &http.Transport{Dial: dial}
	c.Transport = tsp
	return gout.New(c).Debug(false)
}
