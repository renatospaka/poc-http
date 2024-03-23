package main

import (
	"net/http"
)

func midd(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Meu middleware\n"))
}

func main(){
	// log := slog.Default()

	mux := http.NewServeMux()
	mux.HandleFunc("/", midd)
	// mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request){
	// 	w.Write([]byte("Hello, World!\n"))
	// 	fmt.Fprint(w, "caminho raiz (root path)\n")
	// })

	// mux.HandleFunc("GET /task/", func(w http.ResponseWriter, r *http.Request){
	// 	fmt.Fprintf(w, "caminho capturado: %v\n", r.URL.Path)
	// })

	// mux.HandleFunc("GET /task/{id}", func(w http.ResponseWriter, r *http.Request) {
	// 	id := r.PathValue("id")
	// 	fmt.Fprintf(w, "task com id = %v\n", id)
	// })

	// log.Info("Servidor iniciado na porta :8080")
	http.ListenAndServe(":8080", mux)
}
