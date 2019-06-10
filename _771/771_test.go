package _771

import "testing"

func TestNumJewelsInStones1(t *testing.T) {
	J := "aA"
	S := "aAAbbbb"

	stones := numJewelsInStones(J, S)
	if stones != 3 {
		t.Errorf("result:%d, need:%d", stones, 3)
	}
}


func TestNumJewelsInStones2(t *testing.T) {
	J := "z"
	S := "ZZ"

	stones := numJewelsInStones(J, S)
	if stones != 0 {
		t.Errorf("result:%d, need:%d", stones, 0)
	}
}
