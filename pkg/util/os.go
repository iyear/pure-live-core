package util

import (
	"fmt"
	"os"
)

// FileExists checks if a file exists
func FileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil || os.IsExist(err)
}

// MemoryHuman ize converts bytes to human readable format
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
