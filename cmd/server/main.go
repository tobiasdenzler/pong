package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

// Start a simple webserver that returns HTTP status 200
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Starting server on :4455")
	err := http.ListenAndServe(":4455", mux)
	log.Fatal(err)
}
