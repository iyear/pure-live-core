package packet

import (
	_ "embed"
	"errors"
	"github.com/dop251/goja"
	"sync"
)

//go:embed decode.js
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

func Decode(data []byte) (string, error) {
	vm := pool.Get().(*goja.Runtime)

	decode, f := goja.AssertFunction(vm.Get("decode"))
	if !f {
		return "", errors.New("can not assert decode function")
	}

	r, err := decode(goja.Undefined(), vm.ToValue(vm.NewArrayBuffer(data)))
	if err != nil {
		return "", err
	}

	pool.Put(vm)
	return r.String(), nil
}
