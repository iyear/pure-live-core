package util

import "os"

func FileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil || os.IsExist(err)
}
