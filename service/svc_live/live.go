package svc_live

import (
	"github.com/iyear/pure-live-core/model"
	"github.com/iyear/pure-live-core/pkg/client"
	"github.com/iyear/pure-live-core/pkg/conf"
)

func GetRoomInfo(plat string, room string) (*model.RoomInfo, error) {
	var (
		c    model.Client
		info *model.RoomInfo
		err  error
	)

	if c, err = client.GetClient(plat); err != nil {
		return nil, err
	}
	if info, err = c.GetRoomInfo(room); err != nil {
		return nil, err
	}
	return info, nil
}

func GetPlayURL(plat string, room string) (*model.PlayURL, error) {
	var (
		cli model.Client
		url *model.PlayURL
		err error
	)

	if cli, err = client.GetClient(plat); err != nil {
		return nil, err
	}
	if url, err = cli.GetPlayURL(room, conf.QnBest); err != nil {
		return nil, err
	}
	return url, nil
}
