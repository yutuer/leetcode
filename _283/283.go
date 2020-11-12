package _283

func moveZeroes(nums []int) {
	zeroIndex := -1
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		if v == 0 {
			if zeroIndex == -1 {
				zeroIndex = i
			}
		} else {
			if (zeroIndex > -1) {
				exChange(nums, zeroIndex, i)
				zeroIndex += 1
			}
		}
	}
}

func exChange(nums []int, zeroIndex, numIndex int) {
	tmp := nums[numIndex]
	nums[numIndex] = 0
	nums[zeroIndex] = tmp
}
