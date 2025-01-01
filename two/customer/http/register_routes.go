package http

import (
	"github.com/renatospaka/poc-http/customer/adapter"
	"github.com/renatospaka/poc-http/server"
)

func RegisterRoutes(mux *server.CustomMux) {
	customers := mux.Group("/customers")
	customers.Route("GET", "", adapter.ListCustomers)
	customers.Route("POST", "", adapter.CreateCustomer)
	customers.Route("PUT", "", adapter.UpdateCustomer)
	customers.Route("DELETE", "", adapter.DeleteCustomer)
}
