package _41_reverseStr

//leetcode submit region begin(Prohibit modification and deletion)
func reverseStr(s string, k int) string {
	ss := []byte(s)
	length := len(s)
	for i := 0; i < length; i += 2 * k {
		if i+k <= length {
			//reverse(ss[i : i+k])
			reverse2(ss, i, i+k-1)
		} else {
			reverse2(ss, i, length-1)
			//reverse(ss[i:length])
		}
	}
	return string(ss)
}

func reverse2(b []byte, left, right int) {
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}

func reverse(b []byte) {
	left := 0
	right := len(b) - 1
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}

//leetcode submit region end(Prohibit modification and deletion)
