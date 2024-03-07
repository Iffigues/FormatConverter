package main

import (
	"errors"
	"fmt"
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// typer return map of extention manage by PKL
func typer(a string) (string, error) {
	e := map[string]string{
		"json":       "json",
		"jsonnet":    "jsonnet",
		"pcf":        "pcf",
		"properties": "properties",
		"plist":      "plist",
		"textproto":  "textproto",
		"xml":        "xml",
		"yaml":       "yaml",
	}
	if val, ok := e[a]; ok {
		return val, nil
	}
	return "", errors.New("ho")
}

// ConvertFiles transform conf file to other conf fil
func ConverteFile(name, path string, types ResponseData, to string) error {
	ee := NewFormat(types.Ct_Label, types.Path)
	if err := ee.CreatePKL("/tmp/file/generatedpkl/" + name + ".pkl"); err != nil {
		return err
	}
	if err := ee.CreateFile(name); err != nil {
		return err
	}

	if err := ee.Exec(name); err != nil {
		return err
	}

	if err := ee.ExecPKL(to); err != nil {
		return err
	}

	if _, err := ee.CreateNew(path + name + "." + to); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// Totype return configuration file type from option of form
func Totype(r *http.Request) ([]string, error) {
	err := r.ParseForm()
	if err != nil {
		return nil, err
	}
	options, ok := r.Form["options"]
	if !ok || len(options) == 0 {
		return nil, fmt.Errorf("no 'options' parameter provided in the request")
	}

	return options, nil
}

// countFormat return number ofconfiguration to transform
func countFormat(typesFile []ResponseData) int {
	i := 0
	prev := ""

	for _, val := range typesFile {
		if val.Ct_Label != prev {
			i = i + 1
			prev = val.Ct_Label
		}
	}
	return i
}

func GetName(e fs.FileInfo) (name string, dir bool) {
	if e.IsDir() {
		return "", true
	}
	i := e.Name()
	filePath := i
	baseName := filepath.Base(filePath)

	// Use filepath.Ext to get the file extension
	ext := filepath.Ext(baseName)

	// Use strings.TrimSuffix to remove the extension from the base name
	nam := strings.TrimSuffix(baseName, ext)
	return nam, false
}

func CreateDirNeeded(path string, types []string) error {
	for _, ty := range types {
		if err := createDir(path + "/" + ty); err != nil {
			return err
		}
	}
	return nil
}

// handler for get file, return path where dowload file or dir
func Converte(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		err := fmt.Errorf("Something went wrong")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	toType, err := Totype(r)
	if err != nil {
		fmt.Println(err)
		err := fmt.Errorf("new error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(toType) == 0 {
		err := fmt.Errorf("new error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	dirName := GetUid().String()
	pathUpload := "/tmp/file/file/" + dirName + "/"
	if err := createDir(pathUpload); err != nil {
		fmt.Println(err)
		return
	}
	if err = uploadFunc(r, pathUpload); err != nil {
		fmt.Println(err)
		err := fmt.Errorf("Something went wrong")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	typesFiles, err := Magika(pathUpload)
	if err != nil {
		fmt.Println(err)
		err := fmt.Errorf("Someting went wrong")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	createTransforme(typesFiles, toType)
}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", Converte).Methods("POST")
	// Create a new CORS handler
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	})
	handler := c.Handler(mux)
	http.Handle("/", handler)
	http.ListenAndServe(":8780", handler)
}
