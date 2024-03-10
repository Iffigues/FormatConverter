// Code generated from Pkl module `converter`. DO NOT EDIT.
package appconfig

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type Converter struct {
	Port string `pkl:"port"`

	GenerateDir string `pkl:"generateDir"`

	GeneratedPKL string `pkl:"generatedPKL"`

	PKLCommand string `pkl:"PKLCommand"`

	PKLTemplate string `pkl:"PKLTemplate"`

	FileDir string `pkl:"fileDir"`

	PythonMagickaAPIUrl string `pkl:"pythonMagickaAPIUrl"`

	LogDir string `pkl:"logDir"`

	Format map[string]string `pkl:"format"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Converter
func LoadFromPath(ctx context.Context, path string) (ret *Converter, err error) {
	evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := evaluator.Close()
		if err == nil {
			err = cerr
		}
	}()
	ret, err = Load(ctx, evaluator, pkl.FileSource(path))
	return ret, err
}

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Converter
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*Converter, error) {
	var ret Converter
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
