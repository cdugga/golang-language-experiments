package main

import "fmt"

func testDefer() {
	for i := 0; i < 5; i++{
		defer fmt.Println("num..", i)
	}
}
