package srv_os

import (
	"github.com/iyear/pure-live/model"
	"github.com/iyear/pure-live/pkg/ps"
	"time"
)

func GetSysCPU() (*model.SysCPU, error) {
	per, err := ps.GetSysCPU(25*time.Millisecond, false)
	if err != nil {
		return nil, err
	}
	return &model.SysCPU{Percent: per[0]}, nil
}

func GetSelfCPU() (*model.SelfCPU, error) {
	per, err := ps.GetSelfCPU()
	if err != nil {
		return nil, err
	}
	return &model.SelfCPU{Percent: per}, nil
}
