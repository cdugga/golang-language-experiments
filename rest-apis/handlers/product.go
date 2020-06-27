// Package classification of Product API
//
// Documentation for Product API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"main.go/data"
	"net/http"
)

// A list of products returns in the response
// swagger:response productsResponse
type productsResponse struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// swagger:response noContent
type productNoContent struct {

}

// swagger:parameters deleteProduct
type productIDParameterWrapper struct {
	// The Id of the product to delete from the database
	// in:path
	// required: true
	ID int	 `json:"id"`
}


type Products struct{
	l *log.Logger
}


func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}


type KeyProduct struct{}

func (p Products) MiddleWareProductValidation(next http.Handler) http.Handler{

	return http.HandlerFunc(func(rw http.ResponseWriter, r*http.Request) {
		prod := &data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[Error] deserializing product", err)
			http.Error(rw, "Error reading product", http.StatusBadRequest)
			return
		}

		// validate the product


		err = prod.Validate()
		if err != nil{
			p.l.Println("[Error] validating product", err)
			http.Error(
				rw,
				fmt.Sprintf("Error validating product: %s", err),
				http.StatusBadRequest,
				)
			return
		}

		// add product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		// call the next handler, which can be another middleware in the chain, or the final handler
		next.ServeHTTP(rw,req)
	})

}









//func (p* Products) ServeHTTP(rw http.ResponseWriter, r *http.Request){
//	if r.Method == http.MethodGet {
//		p.GetProducts(rw, r)
//		return
//	}
//	if r.Method == http.MethodPost {
//		p.addProduct(rw, r)
//		return
//	}
//
//	if r.Method == http.MethodPut {
//
//		reg := regexp.MustCompile(`/([0-9]+)`)
//		g := reg.FindAllStringSubmatch(r.URL.Path, -1)
//
//		p.l.Println(len(g))
//		if len(g) != 1 {
//			p.l.Println(rw, "Invalid URI more than one id")
//			http.Error(rw, "Invalid URL", http.StatusBadRequest)
//			return
//		}
//
//		if len(g[0]) != 2 {
//			p.l.Println(rw, "Invalid URI more than one capture group")
//			http.Error(rw, "Invalid URL", http.StatusBadRequest)
//			return
//		}
//
//		idString := g[0][1]
//		id, err := strconv.Atoi(idString)
//		if err!=nil{
//			p.l.Println(rw, "Invalid URI unable to convert to number")
//			http.Error(rw, "Invalid URL", http.StatusBadRequest)
//			return
//		}
//
//		p.updateProducts(id ,rw, r )
//	}
//}