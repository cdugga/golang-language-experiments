package handlers

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"main.go/data"
	"net/http"
	"strconv"
)

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
		// add product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		// call the next handler, which can be another middleware in the chain, or the final handler
		next.ServeHTTP(rw,req)
	})

}

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

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Requests")

	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	p.l.Printf("Prod: %#v", prod)

	data.AddProduct(prod)
}

func (p* Products) GetProducts(rw http.ResponseWriter, r *http.Request){

	p.l.Println("Handle GET Requests")

	lp := data.GetProducts()

	err :=  lp.ToJSON(rw)  //json.Marshal(lp)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
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