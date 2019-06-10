package _804

import "testing"

func Test804_1(t *testing.T) {
	expected := 2

	words := []string{"gin", "zen", "gig", "msg"}
	n := uniqueMorseRepresentations(words)
	if n != expected {
		t.Errorf("expected:%d , but:%d", expected, n)
	}
}
