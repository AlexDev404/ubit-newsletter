// Filename: main.go
package main

import (
	"flag"     // Package flag implements command-line flag parsing.
	"log/slog" // Package log implements a simple logging package.
	"net/http" // Package http provides HTTP client and server implementations.
	"os"       // Package os provides a platform-independent interface to operating system functionality.
)

type application struct {
	logger *slog.Logger
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from UBIT newsletter"))
}

func main() {
	port := flag.String("port", "4000", "Port to run the server on")
	flag.Parse()
	// Create a new HTTP server
	server := http.NewServeMux()
	server.HandleFunc("/", home)

	// Create a new logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	application := &application{logger: logger}

	// Log the port the server is starting on
	application.logger.Info("Starting server on ", "addr", *port)

	err := http.ListenAndServe(":"+*port, server)
	// Log any errors that occur
	logger.Error(err.Error())
	os.Exit(1)
}
