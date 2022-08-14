package _67_twoSum

func twoSum(numbers []int, target int) []int {
	if len(numbers) < 2 {
		return nil
	}

	lo, hi := 0, len(numbers)-1
	for lo < hi {
		if numbers[lo]+numbers[hi] == target {
			return []int{lo + 1, hi + 1}
		} else if numbers[lo]+numbers[hi] > target {
			hi--
		} else {
			lo++
		}
	}
	return nil
}
