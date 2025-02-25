// Filename: main_test.go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func (app *application) TestHomeHandler(t *testing.T) {
	// Set up a new request
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Create a new recorder
	rec := httptest.NewRecorder()
	// Create a new server
	handler := http.HandlerFunc(app.home)
	handler.ServeHTTP(rec, req)
	// Check the status code
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body
	expected := "Hello from UBIT newsletter"
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}

func (app *application) TestAboutHandler(t *testing.T) {
	// Set up a new request
	req, err := http.NewRequest("GET", "/about", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Create a new recorder
	rec := httptest.NewRecorder()
	// Create a new server
	handler := http.HandlerFunc(app.about)
	handler.ServeHTTP(rec, req)
	// Check the status code
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	// Check the response body
	expected := "About UBIT newsletter"
	if rec.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), expected)
	}
}
