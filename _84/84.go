package _84

func largestRectangleArea(heights []int) int {
	total := 0
	total = maxArray(heights)
	return total
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxArray(heights []int) int {
	a, b, i := maxArea(heights)
	if i < 0 {
		return 0
	}

	var max = i

	if a != nil {
		max = maxInt(maxArray(a), max)
	}

	if b != nil {
		max = maxInt(maxArray(b), max)
	}

	return max
}

func maxArea(heights []int) ([]int, []int, int) {
	aLen := len(heights)
	if aLen == 0 {
		return nil, nil, -1
	}

	if aLen == 1 {
		return nil, nil, heights[0]
	}

	var min = int(^uint(0) >> 1)
	var index = -1

	for i, v := range heights {
		if v < min {
			min = v
			index = i
		}
	}

	var max = heights[index] * aLen

	if index > -1 {
		if index == 0 {
			return nil, heights[index+1:], max
		} else if index == aLen-1 {
			return heights[:index], nil, max
		}
		return heights[:index], heights[index+1:], max
	}

	return nil, nil, -1
}
