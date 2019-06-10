package _509

var m = map[int]int{0: 0, 1: 1}

func fib(N int) int {
	if N == 0 {
		return 0
	}

	if N == 1 {
		return 1
	}

	if v2, ok := m[N-2]; ok {
		if v1, ok := m[N-1]; ok {
			v := v2 + v1
			m[N] = v
			return v
		}
	}

	return fib(N-1) + fib(N-2)
}
