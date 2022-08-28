package _17_letterCombinations

func letterCombinations(digits string) []string {
	digitsMap := [10]string{
		"",     // 0
		"",     // 1
		"abc",  // 2
		"def",  // 3
		"ghi",  // 4
		"jkl",  // 5
		"mno",  // 6
		"pqrs", // 7
		"tuv",  // 8
		"wxyz", // 9
	}

	var res []string
	var track []byte
	if len(digits) == 0 {
		return res
	}
	var backtrack func(level int)
	backtrack = func(level int) {
		if level > len(digits) {
			return
		}
		if level == len(digits) {
			tmp := make([]byte, level)
			copy(tmp, track)
			res = append(res, string(tmp))
			return
		}
		i := digits[level] - '0'
		digitStr := digitsMap[i]
		for i := 0; i < len(digitStr); i++ {
			track = append(track, digitStr[i])
			backtrack(level + 1)
			track = track[:len(track)-1]
		}
	}

	backtrack(0)
	return res
}
