package utils

import (
	"os"
)

func MkDirAllRoot(r *os.Root, path string) {}

func CreateRootFile(root  *os.Root, path, dirName string, fileName string) (PathDir string, err error) {
	file, err := root.Create(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()
	return path + "/" + dirName, nil
}