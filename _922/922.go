package _922

func sortArrayByParityII(A []int) []int {
	for i, j := 0, 1; i < len(A)-1 && j < len(A); {
		if A[i]&1 == 1 && A[j]&1 != 1 { //A[i] 奇数, A[j]偶数
			// 交换
			A[i], A[j] = A[j], A[i]
		} else if A[i]&1 == 1 { //A[i] 奇数, A[j] 也为奇数
			j += 2
		} else if A[j]&1 != 1 { //A[j] 偶数 A[j] 也为偶数
			i += 2
		} else {
			i, j = i+2, j+2
		}
	}
	return A
}
