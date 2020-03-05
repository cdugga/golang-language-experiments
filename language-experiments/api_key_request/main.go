package api_key_request

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Response struct {
	Val string
}

func eventsHandler(w http.ResponseWriter, r *http.Request){
	val1 := new(Response)
	w.Header().Set("Content-Type", "application/json")

	endpoint := "https://hyj6umtyma.execute-api.eu-west-1.amazonaws.com/Stage/events"
	resp, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(resp.StatusCode)
		w.WriteHeader(resp.StatusCode)
		return
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(val1)
	fmt.Fprintln(w,"Response code:", resp.StatusCode)
	log.Print(resp.StatusCode)
}

func FetchEvents(){
	r := mux.NewRouter()
	r.HandleFunc("/fetchEvents", eventsHandler).Methods("GET")

	err := http.ListenAndServe("localhost:3002", r) //, GreetingHandler("Mayo 4 GO"))
	if err != nil{
		log.Fatal(err)
	}
}
