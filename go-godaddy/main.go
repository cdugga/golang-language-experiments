package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetAvailableDomains(){
	req, err := http.NewRequest("GET", os.ExpandEnv("https://api.godaddy.com/v1/domains/cloudstarter.org"), nil)

	if err != nil {
		fmt.Printf("%s", err)
	}
	resp, err := http.DefaultClient.Do(req)
	bodyBytes, err2 := ioutil.ReadAll(resp.Body)

	if err2 != nil {
		fmt.Printf("%s", err2)
	}

	bodyString := string(bodyBytes)

	if err != nil {
		fmt.Printf("%s", err)
	}

	defer resp.Body.Close()
	fmt.Println(bodyString)

}

func main(){
	fmt.Print("Domain checker")

	GetAvailableDomains()
}
