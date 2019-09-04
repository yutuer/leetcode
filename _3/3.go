package _3

func lengthOfLongestSubstring(s string) int {
	var max = 0
	var count = 0
	var pre = 0
	var m = make(map[byte]int)

	l := len(s)
	for i := 0; i < l; i++ {
		b := s[i]
		if m[b] == 0 {
			m[b] = i + 1
			count ++
			if count > max {
				max = count
			}
		} else {
			for j := pre; j < m[b]-1; j++ {
				bb := s[j]
				m[bb] = 0
				count --
			}
			pre = m[b]
			m[b] = i + 1
		}
	}
	return max
}
