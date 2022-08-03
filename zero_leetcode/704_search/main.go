package main

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := l + (r-l)>>1
		if target == nums[mid] {
			return mid
		}
		if target > nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
