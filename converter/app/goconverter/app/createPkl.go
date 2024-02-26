package main

import (
	"io/ioutil"
	"fmt"
	"os"
)

func (f *Format) CreatePKL(fileName string) error {
	// Create or open the file for writing
	fmt.Println("hoho = ", fileName)
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Println("yaya = ", f.PKL)
	// Write content to the file
	err = ioutil.WriteFile(fileName, f.PKL, 0644)
	return err
}

func (f *Format) CreateNew(fileName string) (string, error) {
	f.NewPath = fileName
	// Create or open the file for writing
	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Write content to the file
	err = ioutil.WriteFile(fileName, f.RenderPKL, 0644)
	return fileName, err
}
