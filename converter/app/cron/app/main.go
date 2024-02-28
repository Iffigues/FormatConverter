package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func getCreationTime(filename string) (time.Time, error) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return time.Time{}, err
	}

	// Obtenir les informations spécifiques au système d'exploitation
	sysInfo, ok := fileInfo.Sys().(*syscall.Stat_t)
	if !ok {
		return time.Time{}, fmt.Errorf("Impossible d'obtenir les informations spécifiques au système d'exploitation")
	}

	// Convertir le timestamp du système d'exploitation en time.Time
	creationTime := time.Unix(sysInfo.Ctim.Unix())
	return creationTime, nil
}

func deleteFile(directory string){
	currentTime := time.Now()

	dir, err := os.Open(directory)
	if err != nil {
		return
	}
	defer dir.Close()

	files, err := dir.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		path := directory + file.Name()
		t, err := getCreationTime(path)
		if err != nil {
			continue
		}
		duration := currentTime.Sub(t)
		isGreaterThan30Minutes := duration > 30*time.Minute
		if isGreaterThan30Minutes {
			err = os.Remove(path)
			if err != nil {
				continue
			}
		} 
	}
}

func main() {
	deleteFile("/tmp/file/file")
	deleteFile("/tmp/file/generatedpkl")
	deleteFile("/tmp/file/generate")
}
