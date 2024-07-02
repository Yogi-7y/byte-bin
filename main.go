package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Byte bin!"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Print("Server started on localhost:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
