package _561

import "sort"

func arrayPairSum(nums []int) int {
	sort.Ints(nums)

	var total int
	for i := 0; i < len(nums); i += 2 {
		total += nums[i]
	}
	return total
}
