package main

import "net/http"

func main() {
	server := NewServer(":3000")
	server.Handle("/", http.MethodGet, HandleRoot)
	server.Handle("/user", http.MethodPost, server.AddMiddleware(PostRequest, CheckAuth(), CallLogging()))
	server.Handle("/api", http.MethodPost, server.AddMiddleware(HandleHome, CheckAuth(), CallLogging()))
	err := server.Listen()
	if err != nil {
		panic(err)
	}
}
