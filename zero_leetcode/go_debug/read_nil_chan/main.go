package main

import "fmt"

func main() {
	c := make(chan int)
	c = nil
	n, ok := <-c
	fmt.Printf("n=%v,ok=%v\n", n, ok)
}

// run result
//fatal error: all goroutines are asleep - deadlock!
//
//goroutine 1 [chan send (nil chan)]:
//main.main()
///Users/aoqiwu/aoqi/github/zero-go/zero_leetcode/go_debug/read_nil_chan/main.go:6 +0x44
