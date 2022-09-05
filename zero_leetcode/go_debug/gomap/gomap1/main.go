package main

func main() {
	var m map[string]int

	m["test_panic"] = 1

}

// Output:
//panic: assignment to entry in nil map
//
//goroutine 1 [running]:
//main.main()
///Users/aoqiwu/aoqi/github/zero-go/zero_leetcode/go_debug/gomap/gomap1/main.go:6 +0x44
