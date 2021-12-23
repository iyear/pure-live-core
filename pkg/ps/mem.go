package ps

import (
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
	"os"
)

func GetSysMem() (*mem.VirtualMemoryStat, error) {
	return mem.VirtualMemory()
}

func GetSelfMem() (*process.MemoryInfoStat, error) {
	proc, err := process.NewProcess(int32(os.Getpid()))
	if err != nil {
		return nil, err
	}
	m, err := proc.MemoryInfo()
	if err != nil {
		return nil, err
	}

	return m, nil
}
