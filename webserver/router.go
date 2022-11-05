package main

import (
	"net/http"
)

func NewRouter() *Router {
	return &Router{
		make(map[string]map[string]http.HandlerFunc),
	}
}

type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, pathExist, methodExist := r.FindHandler(request.URL.Path, request.Method)
	if !pathExist {
		//WriteHeader es para indicar el status del request
		w.WriteHeader(http.StatusNotFound)
		//el return nos ayuda a romper la funcion si esto no existe el handler
		return
	}
	if !methodExist {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	//handler para enviar objeto w y request
	handler(w, request)
}

func (r *Router) FindHandler(path, method string) (http.HandlerFunc, bool, bool) {
	_, pathExist := r.rules[path]
	handler, methodExist := r.rules[path][method]
	return handler, pathExist, methodExist
}
