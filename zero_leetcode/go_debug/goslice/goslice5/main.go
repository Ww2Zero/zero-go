package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	fmt.Printf("%p,cap=%v,len=%v\n", s, cap(s), len(s))
	f(s)
	fmt.Println(s)
}

func f(s []int) {
	fmt.Printf("%p,cap=%v,len=%v\n", s, cap(s), len(s))
	// i只是一个副本，不能改变s中元素的值
	//for _, i := range s {
	//	i++
	//}

	for i := range s {
		fmt.Printf("%p\n", &i)
		s[i] += 1
	}
}
