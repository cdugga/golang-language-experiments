package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func httpRequest(link string){
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)

	}
	fmt.Println(link, resp.Status)
	io.Copy(os.Stdout, resp.Body)

}

