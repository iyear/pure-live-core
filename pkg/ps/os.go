package ps

import "github.com/shirou/gopsutil/v3/host"

func GetOsInfo() (*host.InfoStat, error) {
	return host.Info()
}
