package handlers

import (
	"github.com/gorilla/mux"
	"main.go/data"
	"net/http"
	"strconv"
)

// swagger:route PUT /products products updateProduct
// Update a products details
//
// responses:
//	201: noContent
//  404: errorResponse
//  422: errorValidation

// Update handles PUT requests to update products
func (p Products) UpdateProducts(rw http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	id , err:= strconv.Atoi(vars["id"])
	if err != nil {
	http.Error(rw, "Unable to convert it", http.StatusBadRequest)
	return
	}

	p.l.Println("Handle PUT Requests", id)

	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	p.l.Printf("Prod: %#v", prod)

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
	http.Error(rw, "Product not found", http.StatusNotFound)
	return
	}
	if err != nil {
	http.Error(rw, "Product not found", http.StatusInternalServerError)
	return
	}
}