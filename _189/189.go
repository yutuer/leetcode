package _189

func rotate(nums []int, k int) {
	l := len(nums)
	index := k % l     // 从右向左数几个
	split := l - index //从左向右数几个

	for i := split; i < split+(l-split)/2; i++ {
		// 颠倒后部分
		swap(nums, i, (l-1)-(i-split))
	}

	for i := 0; i < split/2; i++ {
		// 颠倒前部分
		swap(nums, i, (split-1)-i)
	}

	for i := 0; i < l/2; i++ {
		//再次颠倒全部
		swap(nums, i, (l-1)-i)
	}
}

func swap(nums []int, i1, i2 int) {
	if len(nums) <= 1 {
		return
	}

	tmp := nums[i1]
	nums[i1] = nums[i2]
	nums[i2] = tmp
}
