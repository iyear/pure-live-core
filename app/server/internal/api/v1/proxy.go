package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/iyear/pure-live-core/pkg/request"
	"io"
	"net/http"
)

func Proxy(c *gin.Context) {
	req, err := http.NewRequestWithContext(c, c.Request.Method, c.GetHeader("PL-URL"), c.Request.Body)

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(req.Body)
	req.Header = c.Request.Header

	req.Header.Del("PL-URL")
	resp, err := request.HTTP().Client().Do(req)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	for k := range resp.Header {
		for j := range resp.Header[k] {
			c.Header(k, resp.Header[k][j])
		}
	}

	c.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)
}
