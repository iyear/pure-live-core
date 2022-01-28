package client

import (
	"fmt"
	"github.com/iyear/pure-live-core/model"
	"github.com/iyear/pure-live-core/pkg/client/internal/bilibili"
	"github.com/iyear/pure-live-core/pkg/client/internal/douyu"
	"github.com/iyear/pure-live-core/pkg/client/internal/egame"
	"github.com/iyear/pure-live-core/pkg/client/internal/huya"
	"github.com/iyear/pure-live-core/pkg/client/internal/inke"
	"github.com/iyear/pure-live-core/pkg/conf"
)

func GetClient(plat string) (model.Client, error) {
	switch plat {
	case conf.PlatBiliBili:
		return bilibili.NewBiliBili()
	case conf.PlatHuya:
		return huya.NewHuya()
	case conf.PlatDouyu:
		return douyu.NewDouyu()
	case conf.PlatEGame:
		return egame.NewEGame()
	case conf.PlatInke:
		return inke.NewInke()
	}
	return nil, fmt.Errorf("unsupported live platform")
}
