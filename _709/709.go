package _709

func toLowerCase(str string) string {
	if len(str) == 0 {
		return str
	}

	const diff = 'a' - 'A'
	bs := []byte(str)
	for i, v := range bs {
		if v < 'A' || v > 'z'{
			continue
		}

		if bs[i] < 'a' {
			bs[i] += diff
		}
	}
	return string(bs)
}
