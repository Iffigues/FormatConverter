package pklTransformer

import (
	"converter/data"
	"converter/utils"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"text/template"
)

type PklTransformer struct {
	wg    sync.WaitGroup
	datas []data.ResponseData
	dirId string
	to    []string
}

type Format struct {
	Format     string
	FileFormat string
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
	tmpl, err := template.ParseFiles("pklTemplate/pkltemplate")
	if err != nil {
		return err
	}
	if err := utils.CreateDir("/tmp/file/generatedpkl/" + p.dirId); err != nil {
		return err
	}
	file, err := os.Create("/tmp/file/generatedpkl/" + p.dirId + name + ".pkl")
	if err != nil {
		return err
	}
	defer file.Close()

	err = tmpl.Execute(file, Format{
		Format:     datas.Ct_Label,
		FileFormat: datas.Path,
	})
	return err
}

func (p *PklTransformer) createPKLFile(datas data.ResponseData) error {
	fileInfo, err := os.Stat(datas.Path)
	if err != nil {
		return err
	}
	name, _ := p.GetName(fileInfo)
	cmd := exec.Command("pkl", "eval", "/tmp/file/generatedpkl/"+p.dirId+name+".pkl", "-o", "/tmp/file/newpkl/"+p.dirId+name+".pkl")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("err, la ", err, string(output))
		return err
	}
	return nil
}

func (p *PklTransformer) createToFormatFile(datas data.ResponseData, to string) error {
	fileInfo, err := os.Stat(datas.Path)
	if err != nil {
		return err
	}
	name, _ := p.GetName(fileInfo)
	cmd := exec.Command("pkl", "eval", "-f", to, "/tmp/file/newpkl/"+p.dirId+name+".pkl", "-o", "/tmp/file/generate/"+p.dirId+to+"/"+name+"."+to)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("err, la ", err, string(output))
		return err
	}
	return nil
}

func (p *PklTransformer) createPKLTemplate(datas data.ResponseData) {
	fmt.Println(1, p.TemplatePkl(datas))
	fmt.Println(2, p.createPKLFile(datas))
	for _, to := range p.to {
		fmt.Println(p.createToFormatFile(datas, to))
	}
}

func (p *PklTransformer) Start() (int, string, error) {
	for _, i := range p.datas {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			p.createPKLTemplate(i)
		}()
	}
	p.wg.Wait()
	dir := "/tmp/file/generate/" + p.dirId
	count, err := utils.CountFiles(dir)
	return count, p.dirId, err
}
