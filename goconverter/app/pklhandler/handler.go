package pklhandler

import (
	"converter/conf"
	"converter/router"
)

type PklHandler struct {
	confs conf.Conf
}

func NewPKLHandler(confs conf.Conf) *PklHandler {
	return &PklHandler{
		confs: confs,
	}
}

func (p *PklHandler) GetRoute() []*router.Router {
	l := router.NewMethods("POST", nil)
	return []*router.Router{
		router.NewRouter([]*router.Methods{l}, "/convert/file", p.UploadFile),
		router.NewRouter([]*router.Methods{l}, "/convert/text", p.TextToformat),
		router.NewRouter([]*router.Methods{l}, "/describe/file", p.DescribeFile),
		router.NewRouter([]*router.Methods{l}, "/describe/file", p.DescribeText),
	}
}
