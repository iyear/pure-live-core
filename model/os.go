package model

type OSInfo struct {
	Uptime          uint64 `json:"uptime"`
	OS              string `json:"os"`
	Platform        string `json:"platform"`
	PlatformVersion string `json:"platform_version"`
	KernelVersion   string `json:"kernel_version"`
	KernelArch      string `json:"kernel_arch"`
}

type SysCPU struct {
	Percent float64 `json:"percent"`
}

type SelfCPU struct {
	Percent float64 `json:"percent"`
}

type SysMem struct {
	Total    uint64 `json:"total"`
	TotalStr string `json:"total_str"`
	Avl      uint64 `json:"avl"`
	AvlStr   string `json:"avl_str"`
}
type SelfMem struct {
	Mem    uint64 `json:"mem"`
	MemStr string `json:"mem_str"`
}
