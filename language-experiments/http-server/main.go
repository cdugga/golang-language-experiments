package http_server

import (
	"fmt"
	"log"
	"net/http"
)

type GreetingHandler string

func (g GreetingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, g)
}

func RunServer(){
	err := http.ListenAndServe("localhost:3002", GreetingHandler("Mayo 4 GO"))
	if err != nil{
		log.Fatal(err)
	}
}
