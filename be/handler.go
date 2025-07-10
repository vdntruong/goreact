package main

import (
	"net/http"
	"path/filepath"
)

// api

func handleAPIWelcome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"message": "Hello from Go API!"}`))
	}
}

func handleAPIHealthCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}

// static file

var (
	uiDir   = "./static"
	uiIndex = "index.html"

	fs = http.FileServer(http.Dir(uiDir))
)

func handleStaticFiles() http.Handler {
	return http.StripPrefix("/static/", fs)
}

func handleStaticIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(uiDir, uiIndex))
	}
}
