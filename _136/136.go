package _136

import "sort"

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	sort.Ints(nums)

	lastIndex := len(nums) - 1
	//先检查最后一个是不是
	if nums[lastIndex] != nums[lastIndex-1] {
		return nums[lastIndex]
	}

	// 否则,肯定在中间
	for i := 0; i < len(nums); i += 2 {
		if nums[i] != nums[ i+1] {
			return nums[i]
		}
	}
	return 0
}
