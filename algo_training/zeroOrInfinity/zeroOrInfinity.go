package zeroOrInfinity

import (
	"fmt"
	"math"
)

func ZeroOrInfinity() {

	fmt.Println("Input 2!:", Going(2))
	fmt.Println("Input 3!:", Going(3))
		fmt.Println("Input 5!:", Going(8))
	fmt.Println("Input 6!:", Going(6))
	fmt.Println("Input 7!:", Going(7))
	fmt.Println("Input 8!:", Going(8))
	}

func Going(n int) float64{
	sum, a := 1.0,1.0
	for i := 0; i < n-1; i++ {
		a = a/float64(n-i);
		sum = sum+a;
	}
	return math.Floor(sum*1000000)/1000000;
}

//* Verbose option - needs fixing

//func Going(n int) float64 {
//	var sum float64
//	for i := n; i >= 1; i-- {
//		sum += float64(factorial(i))
//	}
//	return math.Floor((1/float64(factorial(n)) * sum) *1000000 )/1000000
//}

//func factorial(n int) int {
//	if n == 1 { return n}
//	return n*factorial(n-1)
//}
