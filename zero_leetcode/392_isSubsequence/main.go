package main

func isSubsequence(s string, t string) bool {
	li, lj := len(s), len(t)
	i, j := 0, 0
	for i < li && j < lj {
		if s[i] == t[j] {
			i++
			j++
		} else {
			j++
		}
	}
	return i == li
}
