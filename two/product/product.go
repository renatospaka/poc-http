package product

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/renatospaka/customer-processor-service/cmd/test/server"
// )

// // RegisterRoutes configura as rotas para o grupo /products
// func RegisterRoutes(mux *server.CustomMux) {
// 	// Grupo de rotas para /products
// 	products := mux.Group("/products")
// 	products.Route("GET", "/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "List of products\n")
// 		// curl -X GET http://localhost:8080/products/
// 	})

// 	products.Route("POST", "/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Create a new product\n")
// 		// curl -X POST http://localhost:8080/products/
// 	})

// 	products.Route("PUT", "", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Update a product\n")
// 		// curl -X PUT http://localhost:8080/products
// 	})

// 	products.Route("DELETE", "", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Delete a product\n")
// 		// curl -X DELETE http://localhost:8080/products
// 	})
// }
