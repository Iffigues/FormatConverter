package main

import (
	"os"
	"text/template"
	"fmt"
)

type Format struct {
	id	  string
	pklDirname string
	Format     string
	FileFormat string
	PKLPath    string
	NewPath    string
}

func NewFormat(format, fileFormat string) (*Format, error) {
	dirname := GetUid().String()
	if err := createDir("/tmp/file/generatedpkl/" + dirname); err != nil {
		return nil, err
	}
	return &Format{
		pklDirname: dirname,
		id:         GetUid().String(),
		Format:     format,
		FileFormat: fileFormat,
	}, nil 
}

func (f *Format) CreateFile(name string) error {
	tmpl, err := template.ParseFiles("pklTemplate/pkltemplate")
	if err != nil {
		return err
	}
	file, err := os.Create("/tmp/file/generatedpkl/" + f.pklDirname  + "/" + name + ".pkl")
	if err != nil {
		return err
	}
	defer file.Close()
	err = tmpl.Execute(file, f)
	fmt.Println("err = ",err)
	return err
}
