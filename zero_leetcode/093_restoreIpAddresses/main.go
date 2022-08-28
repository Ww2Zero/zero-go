package _93_restoreIpAddresses

func restoreIpAddresses(s string) []string {
	var res []string
	if len(s) < 4 || len(s) > 12 {
		return res
	}
	var backtrack func(startIndex int, track []string)
	backtrack = func(startIndex int, track []string) {
		//终止条件
		if startIndex == len(s) && len(track) == 4 {
			res = append(res, convertIpStr(track))
		}
		for i := startIndex; i < len(s); i++ {
			//处理
			track = append(track, s[startIndex:i+1])
			if i-startIndex+1 <= 3 && len(track) <= 4 && isValid(s, startIndex, i) {
				//递归
				backtrack(i+1, track)
			} else { //如果首尾超过了3个，或路径多余4个，或前导为0，或大于255，直接回退
				return
			}
			//回溯
			track = track[:len(track)-1]
		}
	}
	backtrack(0, []string{})
	return res
}

func convertIpStr(nums []string) (res string) {
	for i := 0; i < len(nums); i++ {
		res += nums[i]
		if i != len(nums)-1 {
			res += "."
		}
	}
	return res
}
func isValid(s string, start, end int) bool {
	if start > end {
		return false
	}

	if s[start] == '0' && start != end {
		return false
	}
	num := 0

	for i := start; i <= end; i++ {
		if s[i] > '9' || s[i] < '0' {
			return false
		}
		num = num*10 + int(s[i]-'0')
		if num > 255 {
			return false
		}
	}
	return true
}
