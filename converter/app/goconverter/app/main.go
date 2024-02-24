package main

import (
	"github.com/rs/cors"
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
	formdata := r.MultipartForm
	files := formdata.File["files"]
	if len(files) == 0 {
		return fmt.Errorf("Someting went wrong")
	}
	for _, file := range files {
		fmt.Println("1")
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

func createDir(path string) (error) {
	return os.MkdirAll(path, 0755)
}

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

func createTransforme(typesFile []ResponseData, toType []string) (path string, err error) {
	size := countFormat(typesFile)
	dirName := GetUid().String()
	pathUpload := "/tmp/generate/" + dirName + "/"

	if err := createDir(pathUpload); err != nil {
		return "", err
	}
	if size == 1 {
		return pathUpload, nil
	}

	for _, val := range typesFile {
		fmt.Println(val)
	}
	return pathUpload, nil
}

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
	pathUpload :=  "/tmp/file/" + dirName + "/"
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
