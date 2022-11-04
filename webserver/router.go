package main

import (
	"net/http"
)

func NewRouter() *Router {
	return &Router{
		make(map[string]http.HandlerFunc),
	}
}

type Router struct {
	rules map[string]http.HandlerFunc
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, exist := r.FindHandler(request.URL.Path)
	if !exist {
		//WriteHeader es para indicar el status del request
		w.WriteHeader(http.StatusNotFound)
		//el return nos ayuda a romper la funcion si esto no existe el handler
		return
	}
	//handler para enviar objeto w y request
	handler(w, request)
}

func (r *Router) FindHandler(path string) (http.HandlerFunc, bool) {
	handler, exist := r.rules[path]
	return handler, exist
}
