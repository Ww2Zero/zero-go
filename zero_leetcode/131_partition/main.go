package _31_partition

func partition(s string) [][]string {
	var res [][]string
	var track []string
	var backtrack func(start int)
	backtrack = func(start int) {
		if start >= len(s) {
			tmp := make([]string, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}

		for i := start; i < len(s); i++ {
			if isPartition(s, start, i) {
				track = append(track, s[start:i+1])
			} else {
				continue
			}

			backtrack(i + 1)
			track = track[:len(track)-1]
		}

	}
	backtrack(0)
	return res
}

func isPartition(s string, startIndex, endIndex int) bool {
	left := startIndex
	right := endIndex

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}
