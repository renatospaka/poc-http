package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoute(t *testing.T) {
	mux := connectingToServer(t)
	defer mux.Close()

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}

	mux.Route("GET", "/route", handler)

	// mu.RLock() // Lock for reading to ensure the route is added
	assert.Equal(t, 1, len(mux.Rotas()))
	// mu.RUnlock()

	req := httptest.NewRequest("GET", "/route", nil)
	rr := httptest.NewRecorder()

	mux.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "OK", rr.Body.String())

	// Test not found
	req = httptest.NewRequest("GET", "/notfound", nil)
	rr = httptest.NewRecorder()

	mux.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestGroup(t *testing.T) {
	mux := connectingToServer(t)
	defer mux.Close()

	group := mux.Group("/api")
	assert.NotNil(t, group)
	assert.Equal(t, "/api", group.Prefix())
	assert.Equal(t, mux, group.Mux())

	group.Route("GET", "/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Users route"))
	})

	req, err := http.NewRequest("GET", "/api/users", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "Users route", rr.Body.String())

	// mu.RLock()
	assert.Equal(t, 1, len(mux.Rotas())) // Check if the route was added through the group
	route := mux.Rotas()[0]
	assert.Equal(t, "GET", route.Method)
	assert.Equal(t, "/api/users", route.Pattern)
	// mu.RUnlock()

	req, err = http.NewRequest("GET", "/api/notfound", nil)
	assert.NoError(t, err)

	rr = httptest.NewRecorder()
	mux.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusNotFound, rr.Code)
}

// Locking functions for concurrent testing access to the routes slice
// func lockRoutesForRead() {
// 	// mu.RLock()
// }

// func unlockRoutesForRead() {
// 	// mu.RUnlock()
// }
// func lockRoutesForWrite() {
// 	// mu.Lock()
// }

// func unlockRoutesForWrite() {
// 	// mu.Unlock()
// }
