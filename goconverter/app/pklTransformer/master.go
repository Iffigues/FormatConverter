package pklTransformer

import (
	"converter/data"
	"converter/utils"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type PklTransformer struct {
	wg    sync.WaitGroup
	datas []data.ResponseData
	dirId string
	to    []string
}

func NewPklTransformer(datas []data.ResponseData, to []string) *PklTransformer {
	return &PklTransformer{
		datas: datas,
		to:    to,
		dirId: utils.GetUid().String() + "/",
	}
}

func (p *PklTransformer) GetName(e fs.FileInfo) (name string, dir bool) {
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

func (p *PklTransformer) TemplatePkl(datas data.ResponseData) error {
	fileInfo, err := os.Stat(datas.Path)
	if err != nil {
		return err
	}
	name, _ := p.GetName(fileInfo)
	fmt.Println(name)
	/*tmpl, err := template.ParseFiles("pklTemplate/pkltemplate")
	if err != nil {
		return err
	}
	file, err := os.Create("/tmp/file/generatedpkl/" + p.dirId + name + ".pkl")
	*/
	return nil
}

func (p *PklTransformer) create(datas data.ResponseData) {
	p.TemplatePkl(datas)
}

func (p *PklTransformer) Start() {
	for _, i := range p.datas {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			p.create(i)
		}()
	}
	p.wg.Wait()
}
