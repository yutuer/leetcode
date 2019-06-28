package _26

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var t = nums[0]
	for i := 0; i < len(nums); {
		if i == 0 {
			i++
		} else {
			v := nums[i]

			if v == t {
				if i == len(nums)-1 {
					nums = nums[:i]
				} else {
					nums = append(nums[:i], nums[i+1:] ...)
				}
			} else {
				i++
				t = v
			}
		}
	}
	return len(nums)
}
