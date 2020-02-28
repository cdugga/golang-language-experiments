package main

import "fmt"

func arrayAndSliceOps() {
	coinArry := [3]string{"NEO", "ZCASH", "BTC"}
	slice := coinArry[:2]
	fmt.Println(slice)
	fmt.Println("metadata...\n length: ", len(slice), "\n capacity: ", cap(slice))

	slice = append(slice, "ETH")
	fmt.Println("after append..", slice)
}
