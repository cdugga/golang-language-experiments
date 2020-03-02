package http_server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/url"
)

type GreetingHandler string

const devs = `[
	{"Title":"cdugga", "Skill": "full stack developer"},
	{"Title": "gholly", "Skill": "full stack teacher"}
]`

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
	u, _ := url.Parse(r.URL.String())

	params := u.Query()
	searchTerm := params.Get("name")
	fmt.Println("Search Query is: ", searchTerm)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(devs)

}

func RunServer(){
	r := mux.NewRouter()
	r.HandleFunc("/headers", headerHandler).Methods("GET")
	r.HandleFunc("/devteam", handler).Methods("GET")

	err := http.ListenAndServe("localhost:3002", r) //, GreetingHandler("Mayo 4 GO"))

	if err != nil{
		log.Fatal(err)
	}
}
