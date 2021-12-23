package global

import (
	"fmt"
	"runtime"
)

const Version = "v0.1.0"

func GetRuntime() string {
	return fmt.Sprintf("%s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH)
}
