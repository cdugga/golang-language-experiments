package main

import (	
	"fmt"
	"rsc.io/quote"
)

var y = 42
var z = "Shaken, not stirred"

func main() {
	x := 42 // declares and assigns

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)
    quote.Hello() // "Hello, world."
}