package customer_test

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/renatospaka/customer-processor-service/cmd/test/customer"
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

// 	customer.RegisterRoutes(mux)

// 	testCases := []struct {
// 		method, pattern string
// 		expectedBody    string
// 	}{
// 		{"GET", "/customers", "List of customers\n"},
// 		{"POST", "/customers", "Create a new customer\n"},
// 		{"PUT", "/customers", "Update a customer\n"},
// 		{"DELETE", "/customers", "Delete a customer\n"},
// 	}

// 	for _, tc := range testCases {
// 		req, err := http.NewRequest(tc.method, tc.pattern, nil)
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		rr := httptest.NewRecorder()
// 		mux.ServeHTTP(rr, req)

// 		if rr.Body.String() != tc.expectedBody {
// 			t.Errorf("for %s %s: expected body %q, got %q", tc.method, tc.pattern, tc.expectedBody, rr.Body.String())
// 		}
// 	}
// }
