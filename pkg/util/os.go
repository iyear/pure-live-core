package util

import (
	"fmt"
	"os"
)

func FileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil || os.IsExist(err)
}

func MemoryHuman(m uint64) string {
	if m < 1024 {
		return fmt.Sprintf("%d B", m)
	} else if m < 1024*1024 {
		return fmt.Sprintf("%.2f KB", float64(m)/1024)
	} else if m < 1024*1024*1024 {
		return fmt.Sprintf("%.2f MB", float64(m)/1024/1024)
	} else {
		return fmt.Sprintf("%.2f GB", float64(m)/1024/1024/1024)
	}
}
