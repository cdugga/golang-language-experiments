package main

import "fmt"

func testForLoop() {
	colors := map[string]string{
		"red" : "#fffbbb",
		"green": "#f12345",
		"white": "#ffffff",
	}

	// add new color and verify it exist
	colors["blue"] = "#876541";
	elem, ok := colors["blue"]
	if ok {
		fmt.Println("element lookup result.. " , elem)
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

