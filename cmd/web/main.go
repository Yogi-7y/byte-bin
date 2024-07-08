package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/byte/view", byteView)
	mux.HandleFunc("/byte/create", byteCreate)

	log.Print("Server started on localhost:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
