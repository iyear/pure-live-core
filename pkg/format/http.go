package format

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iyear/pure-live-core/pkg/ecode"
	"net/http"
)

func HTTP(c *gin.Context, code int, err error, data interface{}) {
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
