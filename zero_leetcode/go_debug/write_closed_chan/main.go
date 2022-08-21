package main

func main() {
	c := make(chan int)
	close(c)
	c <- 1
}

// run result
// panic: send on closed channel
//
//goroutine 1 [running]:
//main.main()
//        /Users/aoqiwu/aoqi/github/zero-go/zero_leetcode/go_debug/read_close_chan/main.go:6 +0x58
