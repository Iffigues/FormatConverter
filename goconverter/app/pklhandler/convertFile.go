package pklhandler

import (
	"converter/data"
	"converter/pklTransformer"
	"converter/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Response struct {
	Count int    `json:"count,omitempty"`
	Path  string `json:"path,omitempty"`
	Error error  `json:"error,omitempty"`
}

func (p *PklHandler) beginConvertion(datas []data.ResponseData, to []string) (int, string, error) {
	pkltrans := pklTransformer.NewPklTransformer(datas, to)
	return pkltrans.Start()
}

func (p *PklHandler) Totype(r *http.Request) ([]string, error) {
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

func (p *PklHandler) UploadFile(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		err := fmt.Errorf("something went wrong")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	toType, err := p.Totype(r)

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

	dirName := utils.GetUid().String()
	pathUpload := p.confs.Cfg.FileDir + dirName + "/"

	if err := utils.CreateDir(pathUpload); err != nil {
		fmt.Println(err)
		return
	}

	if err = utils.UploadFunc(r, pathUpload); err != nil {
		err := fmt.Errorf("something went wrong")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	typesFiles, err := data.Magika(pathUpload)
	if err != nil {
		err := fmt.Errorf("someting went wrong")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	count, path, err := p.beginConvertion(typesFiles, toType)
	if err != nil {
		err := fmt.Errorf("someting went wrong")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res := Response{
		Count: count,
		Path:  strings.TrimRight(path, "/"),
		Error: err,
	}
	jsonData, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if count == 0 {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonData))
}
