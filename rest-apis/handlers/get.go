package handlers

import (
	"main.go/data"
	"net/http"
)


// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
//	200: productsResponse

// GetProducts returns the products from the data store
func (p* Products) GetProducts(rw http.ResponseWriter, r *http.Request){

	p.l.Println("Handle GET Requests")

	lp := data.GetProducts()

	err :=  lp.ToJSON(rw)  //json.Marshal(lp)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}