package tribonacciSequence

import "fmt"

func TestTribonacci(){

	fmt.Print(Tribonacci([3]float64{1, 1, 1}, 10))
}

var totals []float64

func Tribonacci(signature [3]float64, n int) []float64 {
	if n == 0{
		return []float64{}
	}
	var copy []float64
	for _, value := range signature{
		copy = append(copy, value)
	}
	result := sum(copy, n)
	return result
}

func sum(trib []float64, n int)[]float64{

	if n ==0 {
		return trib
	}
	var latestValue float64
	for _, num := range trib {
		latestValue +=num
	}
	newSignature := append(trib, latestValue)

	fmt.Println("New signature: " , newSignature)
	return sum(newSignature, n - 1)
}
