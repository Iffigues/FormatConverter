package main

import (
	"os"
	"text/template"
)

type Format struct {
	id 	string
	Format	string
	FileFormat string
	PKLPath	string
	NewPath string
	PKL []byte
}

func NewFormat(format, fileFormat string) (*Format) {
	return &Format{
		id: "ide",
		Format: format,
		FileFormat: fileFormat,
		PKLPath: "./tmp/file/file.pkl",
	}
}


func (f *Format)CreateFile() (error) {
	tmpl, err := template.ParseFiles("pklTemplate/pkltemplate")
	if err != nil {
		return err
	}

	file, err := os.Create("./tmp/pkl/" + f.id + ".pkl")
	if err != nil {
		return err
	}
	defer file.Close()

	err = tmpl.Execute(file, f)
	return err
}
