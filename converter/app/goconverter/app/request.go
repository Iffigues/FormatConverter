package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ResponseData struct {
	Path string `json:"path"`
	Ct_Label string `json:"ct_label"`
}



func Magika(path string) ([]ResponseData, error) {
	url := "http://gopiko.fr:5000/path/"
	
	response, err := http.Get(url + path)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var responseData []ResponseData
	err = json.Unmarshal(body, &responseData)
	return responseData, err
}
