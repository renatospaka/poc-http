package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /path/", Logger(func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "path capturado\n")
	}))

	mux.HandleFunc("/task/{id}/{$}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "task com id=%v\n", id)
	})

	slog.Info("serviço iniciado na porta :8090")
	http.ListenAndServe("localhost:8090", mux)
}

func Logger(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		handler.ServeHTTP(w, r)

		elapsedTime := time.Since(startTime)

		slog.Info("http request", slog.String("method", r.Method), slog.String("path", r.URL.Path), slog.String("duration", elapsedTime.String()))
	}
}
