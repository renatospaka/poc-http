package adapter

import (
	"fmt"
	"net/http"
)

func ListProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List of products\n")
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create a new product\n")
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update a product\n")
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete a product\n")
}
