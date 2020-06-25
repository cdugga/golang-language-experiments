package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {


	fmt.Print("Hello Rest APIs")

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("GoodBye World")
	})

	http.HandleFunc("/c", func(rw http.ResponseWriter,  r*http.Request) {
		log.Println("Hello World")

		d, err := ioutil.ReadAll(r.Body)
		if err != nil {

			http.Error(rw, "Oooops", http.StatusBadRequest)
			return
		}

		log.Printf("Data %s\n", d)

		fmt.Fprintf(rw, "Hello %s", d)
	})



	http.ListenAndServe(":8080", nil)
}
