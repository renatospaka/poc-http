package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /path/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "path capturado %v\n", r.URL.Path)
	})

	mux.HandleFunc("GET /task/{id}/{$}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "task com id=%v\n", id)
	})
	slog.Info("serviço iniciado na porta :8090")
	http.ListenAndServe("localhost:8090", mux)
}
