package main

func firstBadVersion(n int) int {
	l, r := 1, n
	for l <= r {
		m := l + (r-l)>>1
		if isBadVersion(m) {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}

const StartBadVersion = 4

func isBadVersion(n int) bool {
	return n >= StartBadVersion
}
