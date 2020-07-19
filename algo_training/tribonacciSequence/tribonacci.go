package tribonacciSequence

import "fmt"
//Expect(Tribonacci([3]float64{1, 1, 1}, 10)).To(Equal([]float64{1, 1, 1, 3, 5, 9, 17, 31, 57, 105}))
//Expect(Tribonacci([3]float64{0, 0, 1}, 10)).To(Equal([]float64{0, 0, 1, 1, 2, 4, 7, 13, 24, 44}))
//Expect(Tribonacci([3]float64{0, 1, 1}, 10)).To(Equal([]float64{0, 1, 1, 2, 4, 7, 13, 24, 44, 81}))
//Expect(Tribonacci([3]float64{1, 0, 0}, 10)).To(Equal([]float64{1, 0, 0, 1, 1, 2, 4, 7, 13, 24}))
//Expect(Tribonacci([3]float64{0, 0, 0}, 10)).To(Equal([]float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}))
//Expect(Tribonacci([3]float64{1, 2, 3}, 10)).To(Equal([]float64{1, 2, 3, 6, 11, 20, 37, 68, 125, 230}))
//Expect(Tribonacci([3]float64{3, 2, 1}, 10)).To(Equal([]float64{3, 2, 1, 6, 9, 16, 31, 56, 103, 190}))
//Expect(Tribonacci([3]float64{1, 1, 1}, 1)).To(Equal([]float64{1}))
//Expect(Tribonacci([3]float64{300, 200, 100}, 0)).To(Equal([]float64{}))
//Expect(Tribonacci([3]float64{0.5, 0.5, 0.5}, 30)).To(Equal([]float64{0.5, 0.5, 0.5, 1.5, 2.5, 4.5, 8.5, 15.5, 28.5, 52.5, 96.5, 177.5, 326.5, 600.5, 1104.5, 2031.5, 3736.5, 6872.5, 12640.5, 23249.5, 42762.5, 78652.5, 144664.5, 266079.5, 489396.5, 900140.5, 1655616.5, 3045153.5, 5600910.5, 10301680.5}))

func TestTribonacci(){

	fmt.Print(Tribonacci([3]float64{3, 2, 1}, 10))
	fmt.Print(Tribonacci([3]float64{1, 1, 1}, 1))

}

func Tribonacci(signature [3]float64, n int) []float64 {
	if n == 0{return []float64{}
	} else if n < 3 { return signature[:n]}
	return sum(signature[:], n-3)
}

func sum(trib []float64, n int)[]float64{
	if n == 0 {return trib}
	var latestValue float64
	for _, num := range trib[len(trib)-3:] {
		latestValue +=num
	}
	return sum(append(trib, latestValue), n - 1)
}
