package main

import (
	"fmt"
	"strings"
)

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
