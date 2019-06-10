package _905

func sortArrayByParity(A []int) []int {
	var changePos = -1

	aLen := len(A)
	for i := 0; i < aLen; i++ {
		if A[i]&1 != 1 { //偶数
			if changePos > -1 {
				swap(A, i, changePos)
				changePos++
			}
		} else {
			if changePos == -1 {
				changePos = i
			}
		}
	}
	return A
}

func swap(A []int, i, j int) {
	tmp := A[i]
	A[i] = A[j]
	A[j] = tmp
}

func sortArrayByParity2(A []int) []int {
	var leftCorrect, rightCorrect = true, true

	aLen := len(A)
	for i, j := 0, aLen-1; i < j; {
		if leftCorrect {
			if A[i]&1 == 1 { //奇数
				leftCorrect = false
			} else {
				i ++
			}
		}

		if rightCorrect {
			if A[j]&1 != 1 { //偶数
				rightCorrect = false
			} else {
				j--
			}
		}

		if !leftCorrect && !rightCorrect {
			swap(A, i, j)

			leftCorrect, rightCorrect = true, true
		}
	}
	return A
}
