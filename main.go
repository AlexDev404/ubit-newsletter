// FIlename: main.go
package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from UBIT newsletter"))
}

func main() {
	// Create a new HTTP server
	server := http.NewServeMux()
	server.HandleFunc("/", home)
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", server)
	log.Fatal(err)
}
