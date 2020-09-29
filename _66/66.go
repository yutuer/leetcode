package _66

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + 1
		if sum >= 10 {
			digits[i] = 0
			if (i == 0) {
				digits = append([]int{1}, digits...)
			}
		} else {
			digits[i] = sum
			break
		}
	}
	return digits
}
