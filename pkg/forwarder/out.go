package forwarder

import (
	"github.com/q191201771/lal/pkg/base"
	"github.com/q191201771/lal/pkg/httpflv"
	"net"
	"time"
)

func Pull(in In, pullURL string, fn func(tag httpflv.Tag)) error {
	return in.Pull(pullURL, fn)
}

func OutLoop(conn net.Conn, pullURL string, rawURL string, in In) error {
	urlCtx, err := base.ParseHttpUrl(rawURL, "")
	if err != nil {
		return err
	}
	sub := httpflv.NewSubSession(conn, urlCtx, false, "")

	sub.WriteHttpResponseHeader()
	sub.WriteFlvHeader()

	if err = Pull(in, pullURL, func(tag httpflv.Tag) {
		sub.Write(tag.Raw)
	}); err != nil {
		return err
	}

	go func() {
		_ = sub.RunLoop()
	}()

	tick := time.NewTicker(500 * time.Millisecond)
	for {
		<-tick.C
		_, write := sub.IsAlive()
		if !write {
			break
		}
	}
	tick.Stop()

	if err = sub.Dispose(); err != nil {
		return err
	}
	if err = in.Shutdown(); err != nil {
		return err
	}
	return nil
}
