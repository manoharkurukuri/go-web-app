// Test the main function

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain(t *testing.T) {
	req, err := http.NewRequest("GET", "/home", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(homePage)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Just verify the code not html content
	expected := "text/html; charset=utf-8"
	if contentType := rr.Header().Get("Content-Type"); contentType != expected {
		t.Errorf("handler returned unexpected content type: got %v want %v",
			contentType, expected)
	}
}

func TestRoutes(t *testing.T) {
	handler := newHandler()
	tests := []struct {
		path     string
		status   int
		location string
	}{
		{"/", http.StatusFound, "/home"},
		{"/home", http.StatusOK, ""},
		{"/courses", http.StatusOK, ""},
		{"/about", http.StatusOK, ""},
		{"/contact", http.StatusOK, ""},
		{"/missing", http.StatusNotFound, ""},
	}
	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, httptest.NewRequest(http.MethodGet, tt.path, nil))
			if rr.Code != tt.status {
				t.Errorf("status = %d, want %d", rr.Code, tt.status)
			}
			if location := rr.Header().Get("Location"); location != tt.location {
				t.Errorf("Location = %q, want %q", location, tt.location)
			}
			if tt.status == http.StatusOK {
				if contentType := rr.Header().Get("Content-Type"); contentType != "text/html; charset=utf-8" {
					t.Errorf("unexpected Content-Type: %q", contentType)
				}
			}
		})
	}
}
