package _852

func peakIndexInMountainArray(A []int) int {
	aLen := len(A)
	for i := 0; i < aLen-1; i++ {
		if A[i] > A[i+1] {
			return i
		}
	}
	return aLen - 1
}
