package _461

func hammingDistance(x int, y int) int {
	z := x ^ y
	var re int
	for {
		if z == 0 {
			break
		}

		zz := z & 1
		if zz == 1 {
			re ++
		}

		z = z >> 1
	}
	return re
}
