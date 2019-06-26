package _349

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}

	var ret []int
	var m map[int]int = make(map[int]int, 0)

	for _, v := range nums1 {
		m[v] = 1
	}

	for _, v := range nums2 {
		if m[v] == 1 {
			ret = append(ret, v)
			m[v] += 1
		}
	}
	return ret
}
