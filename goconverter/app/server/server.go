package server

import (
	"converter/conf"
	"converter/logger"
	"converter/router"
	"net/http"
	"strings"

	"github.com/rs/cors"
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

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Logic to handle incoming requests using s.handle and i.GetRouter
	for _, i := range s.handle {
		for _, h := range i.GetRouter() {
			s := strings.Split(h, " ")
			if r.URL.Path == s[1] {
				i.F(w, r)
				return
			}
		}
	}
}

func (s *Server) Serve() error {

	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})
	return http.ListenAndServe(s.conf.Cfg.Port, corsHandler.Handler(s))
}
