package data

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"converter/conf"
)

type ResponseData struct {
	Path     string `json:"path"`
	Ct_Label string `json:"ct_label"`
}

func Magika(path string) ([]ResponseData, error) {
	confs, err := conf.NewConf(context.Background(), "pkl/local/appConfig.pkl")
	if err != nil {
		return nil, err
	}
	url := confs.Cfg.PythonMagickaAPIUrl

	response, err := http.Get(url + path)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var responseData []ResponseData
	err = json.Unmarshal(body, &responseData)
	return responseData, err
}

func MagikaDescribeText(str string) ([]ResponseData, error) {
	/*confs, err := conf.NewConf(context.Background(), "pkl/local/appConfig.pkl")
	if err != nil {
		return nil, err
	}
	url := confs.Cfg.PythonMagickaAPIUrl
	*/
	jsonData := map[string]string{"data": str}
	jsonValue, _ := json.Marshal(jsonData)

	response, err := http.Post("http://gopiko.fr:5000/submit", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var responseData []ResponseData
	err = json.Unmarshal(body, &responseData)
	return responseData, err
}
