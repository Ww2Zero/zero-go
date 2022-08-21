package _15_findKthLargest

import (
	"math/rand"
	"time"
)

func findKthLargest(nums []int, k int) int {
	if len(nums) < k {
		return -1
	}
	lo, hi := 0, len(nums)-1
	target := len(nums) - k
	rand.Seed(time.Now().UnixNano())
	for lo <= hi {
		p := partition(nums, lo, hi)
		if p < target {
			lo = p + 1
		} else if p > target {
			hi = p - 1
		} else {
			return nums[p]
		}
	}

	return -1
}
func partition(nums []int, lo, hi int) int {
	picked := rand.Int()%(hi-lo+1) + lo
	nums[picked], nums[lo] = nums[lo], nums[picked]

	pivot := nums[lo]

	i := lo + 1
	j := hi

	for i <= j {
		for i <= j && nums[i] <= pivot {
			i++
		}
		for i <= j && nums[j] > pivot {
			j--
		}
		if i >= j {
			break
		}
		swap(nums, i, j)
	}
	swap(nums, lo, j)
	return j
}

func swap(nums []int, lo, hi int) {
	nums[lo], nums[hi] = nums[hi], nums[lo]
}
