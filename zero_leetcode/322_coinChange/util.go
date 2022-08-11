package _22_coinChange

import "fmt"

func helpN(n int, format string, v ...interface{}) {
	for i := 0; i < n; i++ {
		fmt.Printf("=")
	}
	fmt.Printf(format, v...)
	fmt.Println()
}

func helpP1(n int, format string, v ...interface{}) {
	helpP2(n, n, format, v...)
}
func helpP2(n1, n2 int, format string, v ...interface{}) {
	fmt.Printf("\"node%v\" -> \"node%v\": ", n1, n2)
	fmt.Printf(format, v...)
	fmt.Println()
}
