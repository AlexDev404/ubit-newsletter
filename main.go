// Filename: main.go
package main

import (
	"flag"          // Package flag implements command-line flag parsing.
	"fmt"           // Package fmt implements formatted I/O with functions analogous to C's printf and scanf.
	"html/template" // Package template implements data-driven templates for generating textual output.
	"log/slog"      // Package log implements a simple logging package.
	"net/http"      // Package http provides HTTP client and server implementations.
	"os"            // Package os provides a platform-independent interface to operating system functionality.
)

type application struct {
	logger *slog.Logger
}

type templateData struct {
	Title    string
	Greeting string
}

var data = map[string]string{
	"Title":    "Hello from Go",
	"Greeting": "Hello, World!",
	"Message":  "This is the body of the page",
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("ui/html/home.mustache")
	if err != nil {
		app.logger.Info(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		app.logger.Info(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (app *application) viewTemplate(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len("template/"):]
	fmt.Println("--------- ", path)
	fileName := fmt.Sprintf("%s.%s", path, "mustache")

	// p := templateData{
	// 	Title:    "Go Study",
	// 	Greeting: "Hello World Study Mates",
	// }

	// t, _ := template.ParseFiles(fileName)
	// t.Execute(w, p)
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About UBIT newsletter"))
}

func main() {
	port := flag.String("port", "4000", "Port to run the server on")
	flag.Parse()
	// Create a new HTTP server
	server := http.NewServeMux()
	// Create a new logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	application := &application{logger: logger}

	server.HandleFunc("/", application.home)
	server.HandleFunc("/about", application.about)
	server.HandleFunc("/template/", application.viewTemplate)

	// Log the port the server is starting on
	application.logger.Info("Starting server on ", "addr", *port)

	err := http.ListenAndServe(":"+*port, server)
	// Log any errors that occur
	logger.Error(err.Error())
	os.Exit(1)
}
