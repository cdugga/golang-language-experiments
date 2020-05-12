package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)


const (

)

func GetAvailableDomains(){
	req, err := http.NewRequest("GET", os.ExpandEnv("https://api.godaddy.com/v1/domains/available?domain=${DOMAIN}"), nil)

	if err != nil {
		fmt.Printf("%s", err)
	}

	//req.Header.Set("Authorization", "sso-key ")
	req.Header.Set("Authorization", "sso-key ${GODADDY_API_KEY}:${GODADDY_SECRET_KEY}")

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
	fmt.Println("Domain checker")


	GetAvailableDomains()
}
