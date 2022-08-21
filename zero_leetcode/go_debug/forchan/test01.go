package main

import (
	"fmt"
	"time"
)

const fmtTime = "2006-01-02 15:04:05"

func main() {
	c := make(chan int)
	go func() {
		time.Sleep(1 * time.Second)
		c <- 10
		close(c)
	}()

	for {
		select {
		case x, ok := <-c:
			fmt.Printf("%v，read data from chan  x=%v,ok=%v\n", time.Now().Format(fmtTime), x, ok)
			time.Sleep(500 * time.Millisecond)
			c = nil
		default:
			fmt.Printf("%v，no data from chan \n", time.Now().Format(fmtTime))
			time.Sleep(500 * time.Millisecond)
		}
	}
}
