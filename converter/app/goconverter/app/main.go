package main

import (
	"github.com/rs/cors"
	"os"
	"io"
	"io/fs"
	"errors"
	"net/http"
	"fmt"
	"strings"
	"path/filepath"
	"github.com/gorilla/mux"
)

//typer return map of extention manage by PKL
func typer(a string) (string, error){
	e := map[string]string {
		"json": "json",
		"jsonnet": "jsonnet",
		"pcf":"pcf",
		"properties":"properties",
		"plist":"plist",
		"textproto":"textproto",
		"xml":"xml",
		"yaml":"yaml",
	}
	if val, ok := e[a]; ok {
		return val, nil
	}
	return "", errors.New("ho")
}


// ConvertFiles transform conf file to other conf fil
func ConverteFile(name, path  string, types ResponseData, to string) error {
	fmt.Println(to)
	ee := NewFormat(types.Ct_Label, types.Path)
	err := ee.CreatePKL("/tmp/file/generatedpkl/" + name + ".pkl")
	if err != nil {
		return err
	}
	err = ee.CreateFile(name) 
	if err != nil {
		return err
	}
	err = ee.Exec(name)
	if err != nil {
		return err
	}
	err = ee.ExecPKL(to)
	if err != nil {
		return err
	}
	_, err = ee.CreateNew(path  + name + "." + to)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}


// uploadFunc download one or more files
func uploadFunc(r *http.Request, pathUpload string) (error){
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
		_, err = io.Copy(dst, src)
		if err != nil {
			return err
		}
	}
	return nil
}

// createDir create directories
func createDir(path string) (error) {
	return os.MkdirAll(path, 0755)
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
func countFormat(typesFile []ResponseData) (int) {
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

// create all dir of conf file if multiple file with multiple conf file was send
func CreateGenerated(path string, toType []string) error {
	for _, val := range toType {
		if err := createDir(path + val); err != nil {
			return err
		}
	}
	return nil
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

func OneFormat(path string, typesFile []ResponseData, toType string) (err error){
	for _, val := range typesFile {
		if _, err := typer(val.Ct_Label); err != nil {
			continue
		}
		info, err  := GetInfo(val.Path)
		if err != nil {
			return err
		}
		name, isDir := GetName(info)
		if isDir {
			continue
		}
		if err := ConverteFile(name, path, val, toType); err != nil {
			return err
		}
			
	}
	return nil
}


func CreateDirNeeded(path string, types []string) error {
	for _, ty := range types {
		if err := createDir(path +  "/" + ty); err != nil {
			return err
		}
	}
	return nil
}

func MultipleFormat(path string, typesFile []ResponseData, toType []string) error  {
	if err := CreateDirNeeded(path, toType); err != nil {
		return err
	}
	for _, val := range typesFile {
		if _, err := typer(val.Ct_Label); err != nil {
			continue
		}
		info, err  := GetInfo(val.Path)
		if err != nil {
			return  err
		}
		name, isDir := GetName(info)
		if isDir {
			continue
		}
		for _, tt := range toType {
			if err := ConverteFile(name, path + tt + "/", val, tt); err != nil {
				return err
			}
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

// handler for get file, return path where dowload file or dir
func Converte(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
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
		fmt.Println(err)
		err := fmt.Errorf("new error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return	
	}
	dirName := GetUid().String()
	pathUpload :=  "/tmp/file/file/" + dirName + "/"
	if err :=  createDir(pathUpload); err  != nil {
		fmt.Println(err)
		return
	}
	err = uploadFunc(r, pathUpload)
	if err != nil {
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
     	AllowedOrigins: []string{"*"},
	AllowedMethods: []string{"GET", "POST", "OPTIONS"},
      AllowedHeaders: []string{"*"},
      AllowCredentials: true,
      Debug: true,
   })
	 handler := c.Handler(mux)
	 http.Handle("/", handler)
	http.ListenAndServe(":8780", handler)
}
