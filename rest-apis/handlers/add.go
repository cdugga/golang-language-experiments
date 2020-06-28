package handlers

import (
	"main.go/data"
	"net/http"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// responses:
//	200: productsResponse
//  422: errorValidation
//  501: errorResponse

// Create handles POST requests to add new products
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Requests")

	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	p.l.Printf("Prod: %#v", prod)

	data.AddProduct(prod)
}