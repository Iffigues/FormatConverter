package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// createDir create directories
func CreateDir(path string) error {
	return os.MkdirAll(path, 0755)
}

func CountFiles(dir string) (int, error) {
	var count int
	files, err := os.ReadDir(dir)
	if err != nil {
		return 0, err
	}

	for _, f := range files {
		if f.IsDir() {
			subCount, err := CountFiles(dir + string(os.PathSeparator) + f.Name())
			if err != nil {
				return 0, err
			}
			count += subCount
		} else {
			count++
		}
	}
	return count, nil
}

func GetSingleFileName(dirPath string) (string, error) {
	var fileName string
	foundFile := false

	err := filepath.Walk(dirPath, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !f.IsDir() && !foundFile { // Check for regular file and avoid duplicates
			fileName = path
			foundFile = true
		} else if foundFile {
			return fmt.Errorf("directory contains more than one file")
		}

		return nil
	})

	if err != nil {
		return "", fmt.Errorf("error walking directory: %w", err)
	}

	if !foundFile {
		return "", fmt.Errorf("directory is empty")
	}

	return fileName, nil
}
