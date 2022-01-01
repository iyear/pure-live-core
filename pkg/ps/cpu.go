package ps

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/process"
	"os"
	"time"
)

// GetSysCPU get system cpu info
func GetSysCPU(internal time.Duration, percpu bool) ([]float64, error) {
	return cpu.Percent(internal, percpu)
}

// GetSelfCPU get self cpu info
func GetSelfCPU() (float64, error) {
	proc, err := process.NewProcess(int32(os.Getpid()))
	if err != nil {
		return 0, err
	}
	per, err := proc.CPUPercent()
	if err != nil {
		return 0, err
	}
	return per, nil
}
