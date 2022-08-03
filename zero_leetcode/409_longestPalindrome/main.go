package main

func longestPalindrome(s string) int {
	m := map[byte]int{}
	bytes := []byte(s)
	for _, b := range bytes {
		m[b]++
	}
	res, odd := 0, 0
	for _, v := range m {
		p := v % 2
		res += v - p
		if p == 1 {
			odd = 1
		}
	}
	return res + odd
}
