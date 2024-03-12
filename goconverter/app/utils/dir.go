package utils

import "os"

// createDir create directories
func CreateDir(path string) error {
	return os.MkdirAll(path, 0755)
}
