package _500

import "strings"

var maps = map[rune]int{'a': 2, 'b': 1, 'c': 1, 'd': 2, 'e': 3, 'f': 2, 'g': 2, 'h': 2, 'i': 3, 'j': 2, 'k': 2, 'l': 2, 'm': 1, 'n': 1, 'o': 3, 'p': 3, 'q': 3, 'r': 3, 's': 2, 't': 3, 'u': 3, 'v': 1, 'w': 3, 'x': 1, 'y': 3, 'z': 1}

func findWords(words []string) []string {
	var ret = make([]string, 0)

	for _, v := range words {
		lower := strings.ToLower(v)
		var i = 0
		var is = true
		for _, c := range lower {
			if i == 0 {
				i = maps[c]
			}

			if i != maps[c] {
				is = false
				break
			}
		}

		if is {
			ret = append(ret, v)
		}
	}

	return ret
}
