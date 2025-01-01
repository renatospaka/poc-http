package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/renatospaka/poc-http/server"
	"github.com/stretchr/testify/assert"
)

func connectingToServer(t *testing.T) *server.CustomMux {
	mux, serverReady := server.NewServer(":8081")
	<-serverReady

	assert.NotNil(t, mux)
	assert.IsType(t, &server.CustomMux{}, mux)
	assert.Equal(t, ":8081", mux.Addr())
	return mux
}

func TestNewServer(t *testing.T) {
	mux := connectingToServer(t) // Use a free port
	defer mux.Close()

	assert.NotNil(t, mux)
	assert.NotNil(t, mux.Server())
	assert.Equal(t, ":8081", mux.Addr())
}

func TestServeHTTP(t *testing.T) {
	mux := connectingToServer(t)
	defer mux.Close()

	// Route registration happens after the server is started,
	// so it needs to be done here for these tests.

	mux.Route("GET", "/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Test route"))
	})

	testCases := []struct {
		method, path string
		expectedCode int
		expectedBody string
	}{
		{"GET", "/test", http.StatusOK, "Test route"},
		{"POST", "/test", http.StatusNotFound, ""},    // Method mismatch
		{"GET", "/notfound", http.StatusNotFound, ""}, // Path mismatch
	}

	for _, tc := range testCases {
		req := httptest.NewRequest(tc.method, tc.path, nil)
		rr := httptest.NewRecorder()

		mux.ServeHTTP(rr, req)

		assert.Equal(t, tc.expectedCode, rr.Code)
		if tc.expectedBody != "" {
			assert.Equal(t, tc.expectedBody, rr.Body.String())
		}
	}
}

func TestClose(t *testing.T) {
	mux := connectingToServer(t)
	defer mux.Close()

	err := mux.Close()
	assert.NoError(t, err)

	// Subsequent close should not return an error
	err = mux.Close()
	assert.NoError(t, err)

	muxWithNilServer := &server.CustomMux{}
	err = muxWithNilServer.Close()
	assert.NoError(t, err)
}

func TestAddr(t *testing.T) {
	mux := connectingToServer(t)
	defer mux.Close()

	addr := mux.Addr()
	assert.Equal(t, ":8081", addr)
}
