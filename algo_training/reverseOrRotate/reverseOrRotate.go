package reverseOrRotate

import (
	"fmt"
	"strconv"
	"strings"
)

func TestRevrot(){
	//fmt.Println("Result: ",Revrot("1234", 0))
	//fmt.Println("Result: ",Revrot("", 1))
	//fmt.Println("Result: ",Revrot("", 0))
	//fmt.Println("Result: ",Revrot("1234", 5))
	fmt.Println("Result: ",Revrot("733049910872815764", 5))
	fmt.Println("Result: ",Revrot("66443875", 4))
	fmt.Println("Result: ",Revrot("123456987654", 6))
	fmt.Println("Result: ",Revrot("66443875", 8))
	fmt.Println("Result: ",Revrot("664438769", 8))
	fmt.Println("Result: ",Revrot("123456779", 8))
	fmt.Println("Result: ",Revrot("123456779", 0))
	fmt.Println("Result: ",Revrot("563000655734469485", 4))


}

func Revrot(s string, n int) string {
	if n <= 0 || len(s) == 0 || len(s) < n { return ""}
	intArr := toIntArray(s)
	return splitAndRevrot(intArr, n)
}
func toIntArray(s string) []int {
	arr := strings.Split(s,"")
	var intArr = make([]int, len(arr))
	for i, digit := range arr{
		val, _ := strconv.Atoi(digit)
		intArr[i] = val
	}
	return intArr
}
func splitAndRevrot(digits []int, n int) string{
	numOfChunks:= len(digits)/n
	if len(digits)% n > 0 { numOfChunks = numOfChunks+1 }
	var final string
	pos := 0
	for i := 1; i <= numOfChunks; i++ {
		if pos+n <= len(digits){
			var u = make([]int, n)
			copy(u, digits[pos : pos+n])
			if len(u) <= n {
				final = final + reverseOrRotate(u, cube(u[:]))
				pos += n
			}
		}
	}
	return final
}
func cube(d []int) int{
	var s int
	for _, n := range d { s = s + (n*n*n) }
	return s
}
func reverseOrRotate(d []int, s int) string{
	if s%2 == 0{
		var h []int
		for i := len(d)-1; i >= 0; i-- { h = append(h, d[i]) }
		return format(h)
	}
	return format(append(d[1:], d[0]))
}
func format(n []int) string{
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(n)), ""), "[]")
}