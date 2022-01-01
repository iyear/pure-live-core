package ps

import "github.com/shirou/gopsutil/v3/host"

// GetOsInfo os info
func GetOsInfo() (*host.InfoStat, error) {
	return host.Info()
}
