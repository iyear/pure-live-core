package forwarder

import (
	"github.com/q191201771/lal/pkg/base"
	"github.com/q191201771/lal/pkg/httpflv"
	"github.com/q191201771/lal/pkg/remux"
	"github.com/q191201771/lal/pkg/rtmp"
)

type In interface {
	Pull(pullURL string, fn func(tag httpflv.Tag)) error
}

func GetIn(tp string) In {
	switch tp {
	case "flv":
		return Flv{}
	case "rtmp":
		return Rtmp{}
	default:
		return nil
	}
}

type Flv struct{}

func (s Flv) Pull(pullURL string, fn func(tag httpflv.Tag)) error {
	if err := httpflv.NewPullSession().Pull(pullURL, fn); err != nil {
		return err
	}
	return nil
}

type Rtmp struct{}

func (s Rtmp) Pull(pullURL string, fn func(tag httpflv.Tag)) error {
	if err := rtmp.NewPullSession().Pull(pullURL, func(msg base.RtmpMsg) {
		fn(*remux.RtmpMsg2FlvTag(msg))
	}); err != nil {
		return err
	}
	return nil
}
