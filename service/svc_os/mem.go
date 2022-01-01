package svc_os

import (
	"github.com/iyear/pure-live-core/model"
	"github.com/iyear/pure-live-core/pkg/ps"
	"github.com/iyear/pure-live-core/pkg/util"
)

func GetSysMem() (*model.SysMem, error) {
	m, err := ps.GetSysMem()
	if err != nil {
		return nil, err
	}
	return &model.SysMem{
		Total:    m.Total,
		TotalStr: util.MemoryHuman(m.Total),
		Avl:      m.Available,
		AvlStr:   util.MemoryHuman(m.Available),
	}, nil
}

func GetSelfMem() (*model.SelfMem, error) {
	m, err := ps.GetSelfMem()
	if err != nil {
		return nil, err
	}
	return &model.SelfMem{
		Mem:    m.RSS,
		MemStr: util.MemoryHuman(m.RSS),
	}, nil
}
