package customer

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/renatospaka/customer-processor-service/cmd/test/server"
// )

// // RegisterRoutes configura as rotas para o grupo /customers
// func RegisterRoutes(mux *server.CustomMux) {
// 	// Grupo de rotas para /customers
// 	customers := mux.Group("/customers")
// 	customers.Route("GET", "", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "List of customers\n")
// 		// curl -X GET http://localhost:8080/customers/
// 	})
// 	customers.Route("POST", "", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Create a new customer\n")
// 		// curl -X POST http://localhost:8080/customers/
// 	})
// 	customers.Route("PUT", "", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Update a customer\n")
// 		// curl -X PUT http://localhost:8080/customers
// 	})
// 	customers.Route("DELETE", "", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Delete a customer\n")
// 		// curl -X DELETE http://localhost:8080/customers
// 	})
// }
