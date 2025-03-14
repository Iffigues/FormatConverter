package utils

import (
	"os"
)

func CreateRootDir(path, dirName string, mode os.FileMode) (err error) {
	root, err := os.OpenRoot(path)
	if err != nil {
		return err
	}
	defer root.Close()
	return root.Mkdir(dirName, mode)
}

func CreateRootFile(path, dirName string, fileName string) (PathDir string, err error) {
	root, err := os.OpenRoot(path +  "/" + dirName)
	if err != nil {
		return "", err
	}
	defer root.Close()
	file, err := root.Create(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()
	return path + "/" + dirName, nil
}

func OpenRootFile(path string, fileName string) (file *os.File, err error) {
	root, err := os.OpenRoot(path)
	if err != nil {
		return nil, err
	}
	defer root.Close()
	return root.Open(fileName)
}