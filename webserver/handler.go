package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home")
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the API Endpoint")
}

func PostRequest(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		log.Println("error:", err)
		fmt.Fprintf(w, "error: %v", err)
	}
	response, err := user.ToJson()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("error:", err)
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
