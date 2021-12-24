package forwarder

import (
	"github.com/q191201771/lal/pkg/base"
	"github.com/q191201771/lal/pkg/httpflv"
	"net"
)

func OutLoop(conn net.Conn, pullURL string, rawURL string, in In) error {
	urlCtx, err := base.ParseHttpUrl(rawURL, "")
	if err != nil {
		return err
	}
	sub := httpflv.NewSubSession(conn, urlCtx, false, "")

	sub.WriteHttpResponseHeader()
	sub.WriteFlvHeader()

	if err = Pull(
		in,
		pullURL,
		func(tag httpflv.Tag) {
			sub.Write(tag.Raw)
		},
	); err != nil {
		return err
	}

	if err = sub.RunLoop(); err != nil {
		return err
	}
	return nil
}

func Pull(in In, pullURL string, fn func(tag httpflv.Tag)) error {
	return in.Pull(pullURL, fn)
}
