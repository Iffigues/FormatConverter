package main

import (
	"context"
	"converter/conf"
	"converter/logger"
	"converter/pklhandler"
	"converter/server"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"path/filepath"
	"strings"
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
	ee, err := NewFormat(types.Ct_Label, types.Path)
	if err != nil {
		return err
	}
	if err := ee.CreateFile(name); err != nil {
		return err
	}

	if err := ee.Exec(name, "/tmp/file/generatedpkl/"+ee.pklDirname+"/"+name+".pkl"); err != nil {
		return err
	}

	if err := ee.ExecPKL(to, path+name+"."+to); err != nil {
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
	ext := filepath.Ext(baseName)
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

func Converte(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		err := fmt.Errorf("Something went wrong")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	toType, err := Totype(r)
	if err != nil {
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
		return
	}
	if err = uploadFunc(r, pathUpload); err != nil {
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
	createTransforme(typesFiles, toType)
}

func main() {
	confs, err := conf.NewConf(context.Background(), "pkl/local/appConfig.pkl")
	if err != nil {
		log.Fatal(err)
	}
	logs := logger.NewLog(confs.Cfg.LogDir)

	if err != nil {
		log.Fatal(err)
	}
	serve, err := server.NewServer(confs, logs)
	if err != nil {
		logs.Fatal(err.Error())
		return
	}
	p := pklhandler.NewPKLHandler()
	serve.AddRoute(p)
	serve.Start()
	for {
		serve.Serve()
	}
}
