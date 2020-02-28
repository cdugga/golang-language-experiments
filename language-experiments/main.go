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
	//testForLoop()
	//ifExpr()
	// implicitly builds array then slices it
	coins := []string{"NEO", "ZCASH", "BTC"}
	testSwitch(coins[0])
	//testDefer()
}



func testDefer() {
	for i := 0; i < 5; i++{
		defer fmt.Println("num..", i)
	}
}

func testSwitch(cryptoCoin string) {
	switch coin := cryptoCoin; coin{
		case "NEO":
			fmt.Println("Neo currency")
		case "ZCASH":
			fmt.Println("Zcash ZPK currency")
		case "BTC":
			fmt.Println("Bitcoin currency")
		default:
			fmt.Println("%s.\n", coin)
	}

}


func ifExpr() string{
	// execute short statement before condition
	if v := "COLIN"; v == strings.ToLower("COLIN") {
		fmt.Println(v)
		return v
	} else {
		fmt.Println("no match found for ", v)
	}

	return "no match"
}


func testForLoop() {
	colors := map[string]string{
		"red" : "#fffbbb",
		"green": "#f12345",
		"white": "#ffffff",
	}
	// using both object values
	for color, hex := range colors {
		fmt.Println("Hex for ", color, "is ", hex)
	}
	// ignoring hex value
	for color, _ := range colors {
		fmt.Println("Color:", color)
	}
	// explicitly defining loop boundary
	for i := 0; i <10; i++ {
		fmt.Println("Iterating...", i)
	}
	// using implicit init and post statements
	sum :=1
	for ;  sum < 5; {
		sum += sum
	}
	fmt.Println(sum)
	// above can be replaced with
	for sum < 5 {
		sum +=sum
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
