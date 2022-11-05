package main

import (
	"fmt"
	"net/http"
)

type Server struct {
	port   string
	router *Router
}

func NewServer(portserve string) *Server {
	fmt.Println("Creando Server")
	return &Server{
		port:   portserve,
		router: NewRouter(),
	}
}

func (s *Server) Listen() error {
	fmt.Println("Iniciando server en puerto", s.port)
	http.Handle("/", s.router)
	err := http.ListenAndServe(s.port, nil)
	return err
}

func (s *Server) AddMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, mware := range middlewares {
		f = mware(f)
	}
	return f
}

func (s *Server) Handle(path string, method string, handler http.HandlerFunc) {
	s.router.rules[path] = make(map[string]http.HandlerFunc)
	s.router.rules[path][method] = handler
}
