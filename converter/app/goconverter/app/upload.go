package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// uploadFunc download one or more files
func uploadFunc(r *http.Request, pathUpload string) error {
	formdata := r.MultipartForm
	files := formdata.File["files"]
	if len(files) == 0 {
		return fmt.Errorf("Someting went wrong")
	}
	for _, file := range files {
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
	}
	return nil
}
