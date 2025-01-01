package adapter

import (
	"fmt"
	"net/http"
)

func ListCustomers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List of customers\n")
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create a new customer\n")
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update a customer\n")
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete a customer\n")
}
