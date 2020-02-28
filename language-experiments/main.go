package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type contactAddr struct {
	email string
}

type dev struct {
	name string
	email contactAddr
}

func main() {
	//httpRequest("http://www.google.com")
	//testPointerUpdate()
	testForLoop()
}

func testForLoop() {

	colors := map[string]string{
		"red" : "#fffbbb",
		"green": "#f12345",
		"white": "#ffffff",
	}

	for color, hex := range colors {
		fmt.Println("Hex for ", color, "is ", hex)
	}

	for color, _ := range colors {
		fmt.Println("Color:", color)
	}

}


func testPointerUpdate() {
	d := createDev()
	// will update reference and subsequent requests will use the original memory location
	fmt.Println("**************** UPDATING ONLY REFERENCE TO MEMORY (remain lowercase)*************************")
	fmt.Println(d.toUpper())
	fmt.Println("Original...",d.email.email)
	fmt.Println("Original...",d.name)

	// will update memory location and subsequent requests will use the new updates
	fmt.Println("**************** UPDATING LOCATION IN MEMORY (now all UPPERCASE) *************************")
	d.toUpperWithPointer()
	fmt.Println("Updated...", d.email.email)
	fmt.Println("Updated...", d.name)
}



func createDev() dev {
	d := dev{
		name: "Colin",
		email: contactAddr{
			email: "duggan.colin@gmail.com",
		},
	}
	return d
}

func httpRequest(link string){
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)

	}
	fmt.Println(link, resp.Status)
	io.Copy(os.Stdout, resp.Body)

}

// will update reference to pointer
func (d dev) toUpper() (string, string) {
	return strings.ToUpper(d.name), strings.ToUpper(d.email.email)
}

// will update the point in memory
func (d *dev) toUpperWithPointer() {
	(*d).email.email = strings.ToUpper(d.email.email)
	(*d).name = strings.ToUpper(d.name)
}
