package svc_os

import (
	"github.com/iyear/pure-live-core/model"
	"github.com/iyear/pure-live-core/pkg/ps"
)

func GetOSInfo() (*model.OSInfo, error) {
	info, err := ps.GetOsInfo()
	if err != nil {
		return nil, err
	}

	return &model.OSInfo{
		Uptime:          info.Uptime,
		OS:              info.OS,
		Platform:        info.Platform,
		PlatformVersion: info.PlatformVersion,
		KernelVersion:   info.KernelVersion,
		KernelArch:      info.KernelArch,
	}, nil
}
