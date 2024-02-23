package main

import (
	"os"
	"io"
	"errors"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

func typer(a string) (string, error){
	e := map[string]string {
		"json": "json",
		"jsonnet": "i.jsonnet",
		"pcf":"pcf",
		"properties":"properties",
		"plist":"plist",
		"textproto":"textproto",
		"xml":"xml",
		"yaml":"yml",
	}
	if val, ok := e[a]; ok {
		return val, nil
	}
	return "", errors.New("ho")
}

func ConverteFile(types, path string) error {
	ee := NewFormat("json", "./tmp/file/file.json")
	ee.CreatePKL("./tmp/file/file.pkl")
	ee.ExecPKL("json")
	ee.CreateNew("./tmp/file/test.json")
	ee.Erase()
	return nil
}

func uploadFunc(r *http.Request, pathUpload string) (error){
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		return err
	}

	formdata := r.MultipartForm

	files := formdata.File["files"]
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
		_, err = io.Copy(dst, src)
		return err

	}
	return nil
}

func createDir(path string) (error) {
	return os.MkdirAll(path, 0755)
}

func Totype(r *http.Request) ([]string, error) {
	err := r.ParseForm()
	if err != nil {
		return nil, err
	}
	return r.Form["option"], nil
}

func Converte(w http.ResponseWriter, r *http.Request) {
	dirName := GetUid().String()
	pathUpload :=  "./tmp/file/" + dirName + "/"
	if createDir(pathUpload) != nil {
		return
	}
	err := uploadFunc(r, pathUpload)
	if err != nil {
		err := fmt.Errorf("Something went wrong")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	typesFiles, err := Magika(pathUpload)
	if err != nil {
		err := fmt.Errorf("Someting went wrong")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(typesFiles)
	toType, err := Totype(r)
	if err != nil {
		err := fmt.Errorf("new error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	for _, val := range toType {
		fmt.Println(val)
	}
}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", Converte).Methods("POST")
	http.ListenAndServe(":8780", mux)
}
