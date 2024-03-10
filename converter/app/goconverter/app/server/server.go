package server

import (
	"converter/conf"
	"converter/logger"
	"converter/router"
	"net/http"
)

type GetRoute interface {
	GetRoute() []*router.Router
}

type Server struct {
	conf   conf.Conf
	log    logger.Logs
	route  []GetRoute
	handle []*router.Router
}

func NewServer(conf conf.Conf, log logger.Logs) (*Server, error) {
	return &Server{
		conf: conf,
		log:  log,
	}, nil
}

func (s *Server) AddRoute(r GetRoute) {
	s.route = append(s.route, r)
}

func (s *Server) Start() {
	for _, i := range s.route {
		s.handle = append(s.handle, i.GetRoute()...)
	}
}

func (s *Server) Serve() error {
	for _, i := range s.handle {
		for _, y := range i.GetRouter() {
			http.HandleFunc(y, i.F)
		}
	}
	return http.ListenAndServe(":8889", nil)
}
