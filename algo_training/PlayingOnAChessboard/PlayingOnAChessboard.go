package PlayingOnAChessboard

import (
	"fmt"
)

func Test(){
	fmt.Println(" Input 0:  ", Game(0))
	fmt.Println(" Input 1:  ", Game(1))
	fmt.Println(" Input 7:  ", Game(8))
	fmt.Println(" Input 40:  ", Game(40))
}


func Game(n int) []int {
	s := n*n
	if s % 2 == 0 {
		return []int{s/2}
	} else {
		return []int{s, 2}
	}
}
