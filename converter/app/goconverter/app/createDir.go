package main

import "os"

// createDir create directories
func createDir(path string) error {
	return os.MkdirAll(path, 0755)
}

// create all dir of conf file if multiple file with multiple conf file was send
func CreateGenerated(path string, toType []string) error {
	for _, val := range toType {
		if err := createDir(path + val); err != nil {
			return err
		}
	}
	return nil
}

// create dir for new conf file and start transform
func createTransforme(typesFile []ResponseData, toType []string) (path string, err error) {
	size := len(toType)
	dirName := GetUid().String()
	pathUpload := "/tmp/file/generate/" + dirName + "/"

	if err := createDir(pathUpload); err != nil {
		return "", err
	}
	if size == 1 {
		return pathUpload, OneFormat(pathUpload, typesFile, toType[0])
	}
	return pathUpload, MultipleFormat(pathUpload, typesFile, toType)
}
