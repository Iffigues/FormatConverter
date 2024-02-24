package main

import (
	"fmt"
	"os"
	"io/fs"
	"path/filepath"
)

func GetInfo(filePath string) (fs.FileInfo, error) {
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		fmt.Println("Error getting absolute path:", err)
		return nil, err
	}

	// Get file information
	fileInfo, err := os.Stat(absPath)
	if err != nil {
		return nil, err
	}
	return fileInfo, nil
}

