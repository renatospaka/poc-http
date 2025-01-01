package server

import (
	"net/http"
	"sync"
)

// Rota representa uma rota HTTP com um método e um manipulador
type Rota struct {
	Method  string
	Pattern string
	Handler http.HandlerFunc
}

var (
	mu    sync.RWMutex
	rotas []Rota
)

// Route adiciona uma nova rota com um método HTTP específico
func (c *CustomMux) Route(method, pattern string, handler http.HandlerFunc) {
	mu.Lock()
	defer mu.Unlock()

	rota := Rota{
		Method:  method,
		Pattern: pattern,
		Handler: handler,
	}

	rotas = append(rotas, rota)
}

// Group representa um grupo de rotas com um prefixo comum
type Group struct {
	prefix string
	mux    *CustomMux
}

// Group cria um novo grupo de rotas com um prefixo comum
func (c *CustomMux) Group(prefix string) *Group {
	return &Group{
		prefix: prefix,
		mux:    c,
	}
}

// Route adiciona uma nova rota ao grupo com um método HTTP específico
func (g *Group) Route(method, pattern string, handler http.HandlerFunc) {
	fullPattern := g.prefix + pattern
	g.mux.Route(method, fullPattern, handler)
}

func (g *Group) Prefix() string {
	return g.prefix
}

func (g *Group) Mux() *CustomMux {
	return g.mux
}

func (c *CustomMux) Rotas() []Rota {
	return rotas
}

func (c *CustomMux) Mutex() sync.RWMutex {
	return mu
}
