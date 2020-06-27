package handlers

import (
	"main.go/data"
	"net/http"
)

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Requests")

	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	p.l.Printf("Prod: %#v", prod)

	data.AddProduct(prod)
}