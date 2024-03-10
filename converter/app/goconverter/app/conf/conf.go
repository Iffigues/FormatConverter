package conf

import (
	"context"
	"converter/appconfig"
)

type Conf struct {
	Cfg *appconfig.Converter
}

func NewConf(cont context.Context, path string) (c Conf, err error) {
	cfg, err := appconfig.LoadFromPath(cont, path)
	if err != nil {
		return c, err
	}
	c.Cfg = cfg
	return c, nil
}