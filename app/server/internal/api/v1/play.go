package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iyear/pure-live-core/pkg/forwarder"
	"go.uber.org/zap"
	"net/http"
)

func Play(c *gin.Context) {
	conn, bio, err := c.Writer.Hijack()
	if err != nil {
		zap.S().Warnw("failed to hijack conn", "error", err)
		return
	}

	if bio.Reader.Buffered() != 0 || bio.Writer.Buffered() != 0 {
		zap.S().Warnw("failed to get buffer", "error", err)
		return
	}

	rawUrl := fmt.Sprintf("http://%s%s", c.Request.Host, c.Request.RequestURI)

	in := forwarder.GetIn(c.Query("type"))
	if in == nil {
		c.Status(http.StatusForbidden)
		return
	}

	if err = forwarder.OutLoop(conn, c.Query("url"), rawUrl, in); err != nil {
		return
	}
}
