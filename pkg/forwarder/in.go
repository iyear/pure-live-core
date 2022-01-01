package forwarder

import (
	"github.com/q191201771/lal/pkg/base"
	"github.com/q191201771/lal/pkg/httpflv"
	"github.com/q191201771/lal/pkg/remux"
	"github.com/q191201771/lal/pkg/rtmp"
)

// In live stream pull interface
type In interface {
	Pull(pullURL string, fn func(tag httpflv.Tag)) error
	Shutdown() error
}

// GetIn get pull session
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

// Flv flv pull session
type Flv struct {
	session *httpflv.PullSession
}

// Pull pull flv stream
func (s *Flv) Pull(pullURL string, fn func(tag httpflv.Tag)) error {
	session := httpflv.NewPullSession()
	s.session = session

	if err := session.Pull(pullURL, fn); err != nil {
		return err
	}
	return nil
}

// Shutdown shutdown flv session
func (s *Flv) Shutdown() error {
	return s.session.Dispose()
}

// Rtmp rtmp pull session
type Rtmp struct {
	session *rtmp.PullSession
}

// Pull pull rtmp stream
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

// Shutdown shutdown rtmp session
func (s *Rtmp) Shutdown() error {
	return s.session.Dispose()
}
