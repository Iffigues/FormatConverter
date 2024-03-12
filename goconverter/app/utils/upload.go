package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func UploadFile(pathUpload string, file *multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	dst, err := os.Create(pathUpload + file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	return err
}

// uploadFunc download one or more files
func UploadFunc(r *http.Request, pathUpload string) error {
	formdata := r.MultipartForm
	files := formdata.File["files"]

	if len(files) == 0 {
		return fmt.Errorf("Someting went wrong")
	}
	for _, file := range files {
		if err := UploadFile(pathUpload, file); err != nil {
			return err
		}
	}
	return nil
}
