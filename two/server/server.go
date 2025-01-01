package server

import (
	"fmt"
	"log"
	"net/http"
)

// CustomMux é um wrapper em torno de http.ServeMux que permite adicionar rotas de forma segura
type CustomMux struct {
	mux    *http.ServeMux
	server *http.Server
}

// NewServer cria uma nova instância de CustomMux e inicia o servidor HTTP
func NewServer(addr string) (*CustomMux, chan struct{}) {
	mux := &CustomMux{
		mux: http.NewServeMux(),
	}
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}
	mux.server = server

	// Canal para sinalizar quando o servidor está pronto
	serverReady := make(chan struct{})

	go func() {
		fmt.Println("Starting server on", addr)
		// Sinaliza que o servidor está pronto
		close(serverReady)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("Server failed:", err)
		}
	}()

	return mux, serverReady
}

func (c *CustomMux) Mux() *http.ServeMux {
	return c.mux
}

func (c *CustomMux) Server() *http.Server {
	return c.server
}

func (c *CustomMux) Close() error {
	if c.server == nil {
		return nil
	}
	return c.server.Close()
}

// ServeHTTP implementa o método ServeHTTP para CustomMux
func (c *CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, rota := range rotas {
		log.Printf("Testing all routes: %s | %s\n", r.URL.Path, r.Method)
		if r.URL.Path == rota.Pattern && r.Method == rota.Method {
			rota.Handler(w, r)
			return
		}
	}

	http.NotFound(w, r)
}

func (c *CustomMux) Addr() string {
	return c.server.Addr
}
