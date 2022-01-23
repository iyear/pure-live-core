package inke

import (
	"github.com/iyear/pure-live-core/pkg/conf"
	"testing"
)

func TestInke_GetPlayURL1(t *testing.T) {
	i, _ := NewInke()
	u, err := i.GetPlayURL("297594356", conf.QnBest)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(u)
}

func TestInke_GetRoomInfo(t *testing.T) {
	i, _ := NewInke()
	info, err := i.GetRoomInfo("751044220")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	t.Log(info)
}
