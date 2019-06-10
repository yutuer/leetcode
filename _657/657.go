package _657

var power = map[string]int{"R": 1, "L": -1, "U": 1, "D": -1}
var dir = map[string]int{"R": 0, "L": 0, "U": 1, "D": 1}

func judgeCircle(moves string) bool {
	var values [2]int
	for _, v := range moves {
		vv := string(v)
		values[dir[vv]] += power[vv]
	}

	if values[0] == 0 && values[1] == 0 {
		return true
	}
	return false
}
