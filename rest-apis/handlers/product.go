package handlers

import (
	"log"
	"main.go/data"
	"net/http"
)

type Products struct{
	l *log.Logger
}


func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

func (p* Products) ServeHTTP(rw http.ResponseWriter, r * http.Request){
	if r.Method == http.MethodGet {
		p.GetProducts(rw, r)
		return
	}
	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)

}

func (p* Products) GetProducts(rw http.ResponseWriter, r * http.Request){

	p.l.Println("Handle GET Requests")

	lp := data.GetProducts()

	err :=  lp.ToJSON(rw)  //json.Marshal(lp)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Requests")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json string", http.StatusBadRequest)
	}
	p.l.Printf("Prod: %#v", prod)
}