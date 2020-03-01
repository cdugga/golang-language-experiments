package goroutines

import (
	"fmt"
	"time"
)

func TestGoRoutines() {
	go greeting("GoRoutine World!")
	greeting("Hello")
}

func greeting(s string) {
	fmt.Println(s)
	time.Sleep(100*time.Millisecond)
}
