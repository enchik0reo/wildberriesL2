package server

import (
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func New() *Server {
	s := &Server{}
	s.httpServer = &http.Server{
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return s
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer.Addr = ":" + port
	s.httpServer.Handler = handler
	return s.httpServer.ListenAndServe()
}
