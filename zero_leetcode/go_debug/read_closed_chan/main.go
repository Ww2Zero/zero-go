package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 1
		close(c)
	}()
	time.Sleep(1 * time.Second)
	n, ok := <-c
	fmt.Printf("n=%v,ok=%v\n", n, ok)
	n1, ok1 := <-c
	fmt.Printf("n1=%v,ok1=%v\n", n1, ok1)
	n2, ok2 := <-c
	fmt.Printf("n1=%v,ok1=%v\n", n2, ok2)
}

// run result
//n=1,ok=true
//n1=0,ok1=false
//n1=0,ok1=false
