package product_test

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/renatospaka/customer-processor-service/cmd/test/product"
// 	"github.com/renatospaka/customer-processor-service/cmd/test/server"
// 	"github.com/stretchr/testify/assert"
// )

// func connectingToServer(t *testing.T) *server.CustomMux {
// 	mux, serverReady := server.NewServer(":8080")
// 	<-serverReady

// 	assert.NotNil(t, mux)
// 	assert.IsType(t, &server.CustomMux{}, mux)
// 	return mux
// }

// func TestRegisterRoutes(t *testing.T) {
// 	mux := connectingToServer(t)
// 	defer mux.Close()

// 	product.RegisterRoutes(mux)

// 	testCases := []struct {
// 		method, pattern string
// 		expectedBody    string
// 	}{
// 		{"GET", "/products/", "List of products\n"},
// 		{"POST", "/products/", "Create a new product\n"},
// 		{"PUT", "/products", "Update a product\n"},
// 		{"DELETE", "/products", "Delete a product\n"},
// 	}

// 	for _, tc := range testCases {
// 		req, err := http.NewRequest(tc.method, tc.pattern, nil)
// 		assert.NoError(t, err)

// 		rr := httptest.NewRecorder()
// 		mux.ServeHTTP(rr, req)

// 		assert.Equal(t, tc.expectedBody, rr.Body.String())
// 	}
// }
