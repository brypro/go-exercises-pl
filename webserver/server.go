package main

import "net/http"

type Server struct {
	port   string
	router *Router
}

func NewServer(portserve string) *Server {
	return &Server{
		port:   portserve,
		router: NewRouter(),
	}
}

func (s *Server) Listen() error {
	http.Handle("/", s.router)
	err := http.ListenAndServe(s.port, nil)
	return err
}

func (s *Server) Handle(path string, handler http.HandlerFunc) {
	s.router.rules[path] = handler
}
