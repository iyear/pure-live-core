package client

import (
	"fmt"
	"github.com/iyear/pure-live/conf"
	"github.com/iyear/pure-live/model"
	"github.com/iyear/pure-live/pkg/client/internal/bilibili"
	"github.com/iyear/pure-live/pkg/client/internal/douyu"
	"github.com/iyear/pure-live/pkg/client/internal/huya"
)

func GetClient(plat string) (model.Client, error) {
	switch plat {
	case conf.PlatBiliBili:
		return bilibili.NewBiliBili()
	case conf.PlatHuya:
		return huya.NewHuya()
	case conf.PlatDouyu:
		return douyu.NewDouyu()
	}
	return nil, fmt.Errorf("unsupported live platform")
}
