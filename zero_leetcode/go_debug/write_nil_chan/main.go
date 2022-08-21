package main

func main() {
	c := make(chan int)
	c = nil
	c <- 1
}

// run result
//fatal error: all goroutines are asleep - deadlock!
