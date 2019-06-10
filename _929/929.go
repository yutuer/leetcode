package _929

func numUniqueEmails(emails []string) int {
	m := make(map[string]bool)

	for _, v := range emails {
		email := split(v)
		m[email] = true
	}
	return len(m)
}

func split(s string) string {
	bs := make([]byte, len(s))
	bsIndex := 0
	isSkip := false

	for i, v := range []byte(s) {
		if v == '@' {
			name := string(bs[0:bsIndex])
			return name + string(s[i:])
		} else {
			if v == '+' {
				isSkip = true
			} else {
				if !isSkip {
					if v != '.' {
						bs[bsIndex] = v
						bsIndex++
					}
				}
			}
		}
	}
	return s
}
