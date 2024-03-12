package pklhandler

import (
	"converter/conf"
	"converter/data"
	"converter/pklTransformer"
	"converter/router"
	"converter/utils"
	"fmt"
	"net/http"
)

type PklHandler struct {
	confs conf.Conf
}

func NewPKLHandler(confs conf.Conf) *PklHandler {
	return &PklHandler{
		confs: confs,
	}
}

func (p *PklHandler) beginConvertion(datas []data.ResponseData, path string, to []string) error {
	pkltrans := pklTransformer.NewPklTransformer(datas, to)
	pkltrans.Start()
	return nil
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
		err := fmt.Errorf("Something went wrong")
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
		err := fmt.Errorf("Something went wrong")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	typesFiles, err := data.Magika(pathUpload)
	if err != nil {
		err := fmt.Errorf("Someting went wrong")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	p.beginConvertion(typesFiles, pathUpload, toType)
}

func (p *PklHandler) GetRoute() []*router.Router {
	l := router.NewMethods("POST", nil)
	return []*router.Router{
		router.NewRouter([]*router.Methods{l}, "/", p.UploadFile),
	}
}
