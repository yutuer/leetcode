package _961

func repeatedNTimes(A []int) int {
	m := make(map[int]bool)

	for _, v := range A {
		if m[v] {
			return v
		} else {
			m[v] = true
		}
	}
	return 0
}
