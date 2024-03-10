package pklhandler

import (
	"converter/router"
	"fmt"
	"net/http"
)

type PklHandler struct {
}

func NewPKLHandler() *PklHandler {
	return &PklHandler{}
}

func (p *PklHandler) Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func (p *PklHandler) GetRoute() []*router.Router {
	l := router.NewMethods("GET", nil)
	return []*router.Router{
		router.NewRouter([]*router.Methods{l}, "/", p.Hello),
	}
}
