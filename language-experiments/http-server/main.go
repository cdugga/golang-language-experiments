package http_server

import (
	"fmt"
	"github.com/gorilla/mux"
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

	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr= %q\n", r.RemoteAddr)

	fmt.Fprintf(w, "\n\n\"Accept\" %q", r.Header["Accept"])
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func RunServer(){
	r := mux.NewRouter()

	r.HandleFunc("/headers", headerHandler).Methods("GET")

	http.HandleFunc("/goworld", handler)
	err := http.ListenAndServe("localhost:3002", r) //, GreetingHandler("Mayo 4 GO"))
	if err != nil{
		log.Fatal(err)
	}
}
