package main

import (
	"os"
	"io/ioutil"
)

func (f *Format) CreatePKL() error {
	fileName := f.PKLPath

	// Create or open the file for writing
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write content to the file
	err = ioutil.WriteFile(fileName, f.PKL, 0644)
	return err
}


func (f *Format)CreateNew(i []byte) (string, error) {
	fileName := "./tmp/file/go.json"
	f.NewPath = fileName
	// Create or open the file for writing
	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Write content to the file
	err = ioutil.WriteFile(fileName, i, 0644)
	return fileName, err
}
