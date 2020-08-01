package maxSubArray


func MaximumSubarraySumConcise(numbers []int) int {
	var maxEndingHere, maxSoFar int ;
	for _, s := range numbers{
		maxEndingHere+=s;
		if maxEndingHere < 0 { maxEndingHere = 0 }
		if maxEndingHere > maxSoFar { maxSoFar = maxEndingHere}
	}
	return maxSoFar
}