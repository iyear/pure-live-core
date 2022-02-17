package global

import (
	"fmt"
	"runtime"
)

const (
	// Version pure-live version desc
	Version = "v0.1.2"
)

// GetRuntime get runtime info
func GetRuntime() string {
	return fmt.Sprintf("%s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH)
}
