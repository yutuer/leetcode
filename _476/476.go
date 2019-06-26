package _476

func findComplement(num int) int {
	if num == 0 {
		return 1
	}

	var total = 0xFFFFFFFF
	var i uint = 0
	var t int
	for ; ; {
		t = total - (1<<i - 1)
		if num&t != 0 {
			i++
		} else {
			break
		}
	}
	return (total - t) & ^num
}
