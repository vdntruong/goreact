package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		println("PORT environment variable not set in shell, defaulting to", port)
	}

	mux := http.NewServeMux()

	{
		// API
		mux.Handle("/api/", handleAPIWelcome())
		mux.Handle("/healthz", handleAPIHealthCheck())
		// Static files
		mux.Handle("/static/", handleStaticFiles())
		mux.HandleFunc("/", handleStaticIndex())
	}

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}

	go func() {
		println("Listening on port 8080")
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			println("Error starting server")
			panic(err)
		}
	}()

	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, os.Interrupt)
	<-stopSignal

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		println("Error shutting down server")
		panic(err)
	}

	println("Server stopped")
}
