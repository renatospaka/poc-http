package http

import (
	"github.com/renatospaka/poc-http/product/adapter"
	"github.com/renatospaka/poc-http/server"
)

func RegisterRoutes(mux *server.CustomMux) {
	products := mux.Group("/products")
	products.Route("GET", "", adapter.ListProducts)
	products.Route("POST", "", adapter.CreateProduct)
	products.Route("PUT", "", adapter.UpdateProduct)
	products.Route("DELETE", "", adapter.DeleteProduct)
}
