package format

import "encoding/json"

func WS(tp string, data interface{}) []byte {
	type msg struct {
		Type string      `json:"type"`
		Data interface{} `json:"data,omitempty"`
	}
	b, _ := json.Marshal(&msg{
		Type: tp,
		Data: data,
	})
	return b
}
