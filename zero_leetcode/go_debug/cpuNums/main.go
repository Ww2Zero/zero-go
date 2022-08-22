package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("num=%d", runtime.NumCPU())
}

// Output:
//num=10
