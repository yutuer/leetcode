package _804

func uniqueMorseRepresentations(words []string) int {
	var mores = [26]string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

	m := make(map[string]int)

	for _, v := range words {
		var s string
		for _, w := range v {
			index := w - 'a'
			s += mores[index]
		}
		m[s] += 1
	}
	return len(m)
}
