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

func headerHandler(w http.ResponseWriter, r *http.Request){
	for k,v := range r.Header {
		fmt.Fprintf(w, "Header field %q, Value %q\n", k, v)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func RunServer(){
	http.HandleFunc("/goworld", handler)
	http.HandleFunc("/headers", headerHandler)
	err := http.ListenAndServe("localhost:3002", nil) //, GreetingHandler("Mayo 4 GO"))
	if err != nil{
		log.Fatal(err)
	}
}
