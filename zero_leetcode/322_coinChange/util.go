package _22_coinChange

import "fmt"

func helpN(n int, format string, v ...interface{}) {
	for i := 0; i < n; i++ {
		fmt.Printf("=")
	}
	fmt.Printf(format, v...)
	fmt.Println()
}
