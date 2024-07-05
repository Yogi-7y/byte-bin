package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/byte/view", byteView)
	mux.HandleFunc("/byte/create", byteCreate)

	log.Print("Server started on localhost:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
