package goroutines

import "fmt"

func TestGoRoutines() {

	nums := []int{1, 2, 3, 4, 5, 6}
	// need to create the channel before using
	c := make(chan int)
	// channel is sent to function
	go greeting(nums, c)

	// receive result from channel c
	res := <-c
	fmt.Println(res)
}

func greeting(s []int, c chan int) {
	sum := 0
	for _, v := range s{
		sum +=v
	}
	// result is sent to channel c
	c <- sum
}
