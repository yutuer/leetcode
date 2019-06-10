package _977

import "sort"

func sortedSquares(A []int) []int {
	aLen := len(A)
	if aLen == 1 {
		A[0] = A[0] * A[0]
		return A
	}

	ret := make([]int, aLen)

	var pos = aLen - 1
	for i, j := 0, aLen-1; i <= j; {
		if i == j {
			// 如果走到这种情况, 则为最后一个, 放在第一位
			ret[0] = A[i] * A[i]
			break
		}

		if A[i] <= 0 && A[j] >= 0 {
			as, bs := A[i]*A[i], A[j]*A[j]
			if as > bs {
				ret[pos] = as
				i, pos = i+1, pos-1
			} else if as < bs {
				ret[pos] = bs
				j, pos = j-1, pos-1
			} else {
				ret[pos] = as
				i, pos = i+1, pos-1

				ret[pos] = bs
				j, pos = j-1, pos-1
			}
		} else if A[i] > 0 { // 全部大于0了, 剩余的从头开始一直插入到pos
			for ii := 0; ii <= pos; ii, i = ii+1, i+1 {
				ret[ii] = A[i] * A[i]
			}
			break
		} else if A[j] < 0 {
			for ii := i; ii <= j; ii++ {
				ret[pos] = A[ii] * A[ii]
				pos --
			}
			break
		}
	}
	return ret
}

func sortedSquares2(A []int) []int {
	for i, v := range A {
		A[i] = v * v
	}
	sort.Ints(A)
	return A
}

func sortedSquares3(A []int) []int {
	var ret []int
	for _, a := range A {
		ret = append(ret, a*a)
	}
	sort.Ints(ret)
	return ret
}