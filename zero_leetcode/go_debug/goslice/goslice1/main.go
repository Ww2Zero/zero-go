package main

import "fmt"

func main() {
	s := []int{5}
	s = append(s, 7)
	s = append(s, 9)
	x := append(s, 11)
	y := append(s, 12)
	fmt.Println(s, x, y)
	fmt.Printf("%p,%p,%p", s, x, y)
}

// Output:
//[5 7 9] [5 7 9 12] [5 7 9 12]
//0x1400012c000,0x1400012c000,0x1400012c000
