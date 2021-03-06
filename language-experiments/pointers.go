package main

import (
	"fmt"
	"strings"
)

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

// will update the point in memory
func (d *dev) toUpperWithPointer() {
	(*d).email.email = strings.ToUpper(d.email.email)
	(*d).name = strings.ToUpper(d.name)
}
