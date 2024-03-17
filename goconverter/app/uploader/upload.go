package uploader

import (
	"converter/conf"
	"converter/router"
	"converter/utils"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type UploaderHandler struct {
	confs conf.Conf
}

func NewPUploader(confs conf.Conf) *UploaderHandler {
	return &UploaderHandler{
		confs: confs,
	}
}

func (p *UploaderHandler) downloadHandler(w http.ResponseWriter, path string) {

	// Extract the filename from the request path (adjust as needed)
	fileName := path
	// Open the file
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("eer=", err)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "File notfddfdf found: %s", fileName)
		return
	}
	defer file.Close()
	// Set content headers for download
	w.Header().Set("Content-Type", "application/octet-stream") // Adjust content type if needed
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileName))

	// Copy file content to the response body
	_, err = io.Copy(w, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error downloading file: %v", err)
		return
	}
}

func (p *UploaderHandler) UploadFile(w http.ResponseWriter, r *http.Request) {
	idString := strings.Split(r.URL.Path, "/")[2]
	fmt.Println("/tmp/file/generate/"+idString, r.URL.Path)
	count, err := utils.CountFiles("/tmp/file/generate/" + idString)
	if err != nil {
		fmt.Println(err)
		return
	}
	if count < 1 {
		return
	}
	if count == 1 {
		path, err := utils.GetSingleFileName("/tmp/file/generate/" + idString)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("path=", path)
		p.downloadHandler(w, path)
		return
	}
	if count > 1 {
		l := utils.GetCompType(strings.Split(r.URL.Path, "/"))
		fmt.Println(l)
		utils.ZipDir("/tmp/file/generate/"+idString, "/tmp/file/compressed"+utils.GetUid().String()+".zip")
		return
	}
}

func (p *UploaderHandler) GetRoute() []*router.Router {
	l := router.NewMethods("GET", nil)
	return []*router.Router{
		router.NewRouter([]*router.Methods{l}, "/upload/{id}", p.UploadFile),
		router.NewRouter([]*router.Methods{l}, "/upload/{id}/{compressed}", p.UploadFile),
	}
}
