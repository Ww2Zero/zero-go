package main

import "fmt"

func myAppend(s []int) []int {
	fmt.Printf("%p,cap=%v,len=%v\n", s, cap(s), len(s))

	// 这里 s 虽然改变了，但并不会影响外层函数的 s
	f := append(s, 100)
	fmt.Printf("s=%p\n", s)
	fmt.Printf("f=%p\n", f)
	return f
}

func myAppendPtr(s *[]int) {
	// 会改变外层 s 本身
	*s = append(*s, 100)
	return
}

func main() {
	//s := []int{1, 1, 1}
	s := []int{5}
	s = append(s, 7)
	s = append(s, 9)

	fmt.Printf("%p\n", s)
	newS := myAppend(s)

	fmt.Println(s)
	fmt.Println(newS)
	fmt.Printf("%p,%p\n", s, newS)
	s = newS

	myAppendPtr(&s)
	fmt.Println(s)
}

// Output:
//[1 1 1]
//[1 1 1 100]
//[1 1 1 100 100]
