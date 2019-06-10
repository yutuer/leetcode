package _771

func numJewelsInStones(J string, S string) int {
	const len = 'z' - 'A' + 1
	var bs [len]bool

	for _, v := range J {
		index := v - 'A'
		bs[index] = true
	}

	var total int
	for _, v := range S {
		index := v - 'A'
		if bs[index] {
			total++
		}
	}
	return total
}
