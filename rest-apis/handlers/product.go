package handlers

import (
	"encoding/json"
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

func (p* Products) ServeHTTP(rw http.ResponseWriter, h * http.Request){

		lp := data.GetProducts()

		d, err := json.Marshal(lp)

		if err != nil {
			http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		}
		rw.Write(d)
}