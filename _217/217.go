package _217

import "sort"

func containsDuplicate(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	sort.Ints(nums)
	var v = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == v {
			return true
		}
		v = nums[i]
	}
	return false
}
