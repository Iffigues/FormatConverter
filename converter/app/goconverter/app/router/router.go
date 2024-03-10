package router

import (
	"net/http"
)

type Methods struct {
	Method     string
	MiddleWare []func(http.ResponseWriter, *http.Request)
}

func NewMethods(m string, f []func(http.ResponseWriter, *http.Request)) *Methods {
	return &Methods{
		Method:     m,
		MiddleWare: f,
	}
}

type Router struct {
	Method []*Methods
	Path   string
	F      func(http.ResponseWriter, *http.Request)
}

func NewRouter(method []*Methods, path string, l func(http.ResponseWriter, *http.Request)) *Router {
	return &Router{
		Method: method,
		Path:   path,
		F:      l,
	}
}

func (r *Router) GetRouter() (path []string) {
	for _, val := range r.Method {
		path = append(path, val.Method+" "+r.Path)
	}
	return
}
