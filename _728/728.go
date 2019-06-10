package _728

func selfDividingNumbers(left int, right int) []int {
	ret := make([]int, 0)
	for i := left; i <= right; i++ {
		if (is(i)) {
			ret = append(ret, i)
		}
	}
	return ret
}

func is(i int) bool {
	if i < 10 {
		return true
	}

	for k := i; k > 0; {
		if k%10 == 0 {
			return false
		}

		j := k - (k / 10 * 10)
		if (i%j != 0) {
			return false
		}
		k /= 10
	}
	return true
}
