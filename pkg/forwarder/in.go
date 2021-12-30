package forwarder

import (
	"github.com/q191201771/lal/pkg/base"
	"github.com/q191201771/lal/pkg/httpflv"
	"github.com/q191201771/lal/pkg/remux"
	"github.com/q191201771/lal/pkg/rtmp"
)

type In interface {
	Pull(pullURL string, fn func(tag httpflv.Tag)) error
	Shutdown() error
}

func GetIn(tp string) In {
	switch tp {
	case "flv":
		return &Flv{}
	case "rtmp":
		return &Rtmp{}
	default:
		return nil
	}
}

type Flv struct {
	session *httpflv.PullSession
}

func (s *Flv) Pull(pullURL string, fn func(tag httpflv.Tag)) error {
	session := httpflv.NewPullSession()
	s.session = session

	if err := session.Pull(pullURL, fn); err != nil {
		return err
	}
	return nil
}

func (s *Flv) Shutdown() error {
	return s.session.Dispose()
}

type Rtmp struct {
	session *rtmp.PullSession
}

func (s *Rtmp) Pull(pullURL string, fn func(tag httpflv.Tag)) error {
	session := rtmp.NewPullSession()
	s.session = session

	if err := session.Pull(pullURL, func(msg base.RtmpMsg) {
		fn(*remux.RtmpMsg2FlvTag(msg))
	}); err != nil {
		return err
	}
	return nil
}

func (s *Rtmp) Shutdown() error {
	return s.session.Dispose()
}
