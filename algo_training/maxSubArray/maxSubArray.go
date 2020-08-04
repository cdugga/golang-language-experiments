package maxSubArray

import (
	"fmt"
	"math"
)

func MaxSubArray(){
	fmt.Println("MaximumSubarraySum([]int{})): ",MaximumSubarraySum([]int{}))
	fmt.Println("MaximumSubarraySum([]int{})): ",MaximumSubarraySum([]int{-1,-1,-3,-4,-1}))
	fmt.Println("MaximumSubarraySum([]int{})): ",MaximumSubarraySum([]int{1,3,4}))
	fmt.Println("MaximumSubarraySum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})): ",MaximumSubarraySum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func MaximumSubarraySum(numbers []int) int {
	if len(numbers) == 0 { return 0}
	isNegative := func(number int) bool {return math.Signbit(float64(number))}
	arr := filter(numbers, isNegative)
	if len(arr) == len(numbers){ return 0 }
	if len(arr) == 0 && len(numbers) > 0 { var sum int; for _, s := range numbers { sum +=s }; return sum; }

	var maxEndingHere, maxSoFar int ;
	for _, s := range numbers{
		maxEndingHere+=s;
		if maxEndingHere < 0 { maxEndingHere = 0 }
		if maxEndingHere > maxSoFar { maxSoFar = maxEndingHere}
	}
	return maxSoFar
}

func filter(numbers []int, test func(int)bool)(ret []int){
	for _, number := range numbers { if test(number) { ret = append(ret, number)}}
	return
}