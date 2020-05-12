package tools

import "os"

func FileExists(path string) (isExist bool, isDir bool) {
	s, err := os.Stat(path)
	if err != nil && !os.IsExist(err) {
		isExist = false
		isDir = false
	} else {
		isExist = true
		isDir = s.IsDir()
	}
	return
}
