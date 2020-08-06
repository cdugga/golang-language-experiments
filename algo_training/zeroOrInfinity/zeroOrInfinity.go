package zeroOrInfinity

import (
	"fmt"
)

func ZeroOrInfinity() {

	//fmt.Println("Input 2!:", Going(2))
	//fmt.Println("Input 3!:", Going(3))
	fmt.Println("Input 5!:", Going(5))
	fmt.Println("Input 6!:", Going(6))
	fmt.Println("Input 7!:", Going(7))
	fmt.Println("Input 8!:", Going(8))
}

func Going(n int) float64 {
	var sum float64
	for i := n; i >= 1; i-- {
		sum += float64(factorial(i))
	}
	return 1/float64(factorial(n)) * sum
}

func factorial(n int) int {
	if n == 1 { return n}
	return n*factorial(n-1)
}
