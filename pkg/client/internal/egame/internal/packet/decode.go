package packet

import (
	_ "embed"
	"encoding/json"
	"errors"
	"github.com/dop251/goja"
	"sync"
)

type Resp struct {
	Seq       int `json:"seq"`
	Operation int `json:"operation"`
	Body      []*struct {
		EventId string            `json:"event_id"`
		MsgType int               `json:"msg_type"`
		BinData []json.RawMessage `json:"bin_data"`
	} `json:"body"`
}

//go:embed decode.min.js
var js string

// goja.Runtime 非并发安全
var pool = sync.Pool{
	New: func() interface{} {
		vm := goja.New()
		_ = vm.Set("utf8Decode", vm.ToValue(func(call goja.FunctionCall) goja.Value {
			// 替代 js 中的 TextDecoder，在 golang 中，string 默认为 utf8 格式。直接 string() 强制转换即可实现 bytes -> utf8 string
			buf := call.Argument(0).Export().(goja.ArrayBuffer)
			return vm.ToValue(string(buf.Bytes()))
		}))
		_, _ = vm.RunString(js)
		return vm
	},
}

func Decode(data []byte) (*Resp, error) {
	vm := pool.Get().(*goja.Runtime)

	decode, f := goja.AssertFunction(vm.Get("decode"))
	if !f {
		return nil, errors.New("can not assert decode function")
	}

	r, err := decode(goja.Undefined(), vm.ToValue(vm.NewArrayBuffer(data)))
	if err != nil {
		return nil, err
	}

	var resp Resp
	if err = json.Unmarshal([]byte(r.String()), &resp); err != nil {
		return nil, err
	}

	pool.Put(vm)
	return &resp, nil
}
