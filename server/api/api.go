package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iyear/pure-live/pkg/ecode"
	"net/http"
)

type H map[string]interface{}

func RespFmt(c *gin.Context, code int, err error, data interface{}) {
	type resp struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data,omitempty"`
	}

	var msg = fmt.Sprintf("%s", ecode.GetMsg(code))
	if err != nil {
		msg += fmt.Sprintf(": %s", err.Error())
	}
	c.JSON(http.StatusOK, &resp{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func MsgFmt(tp string, data interface{}) []byte {
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
