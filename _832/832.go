package _832

func flipAndInvertImage(A [][]int) [][]int {
	for _, v := range A {
		do(v)
	}
	return A
}

func do(ints []int) {
	aLen := len(ints)
	if aLen < 1 {
		return
	}

	for i, j := 0, aLen-1; i <= j; {
		// 水平
		if i != j {
			swap(ints, i, j)
		}

		// 反转
		revert(ints, i)
		if i != j {
			revert(ints, j)
		}

		i++
		j--
	}
}

func revert(A []int, i int) {
	A[i] ^= 1
}

func swap(A []int, i, j int) {
	tmp := A[i]
	A[i] = A[j]
	A[j] = tmp
}
