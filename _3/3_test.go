package _3

import "testing"

func TestA(t *testing.T) {
	substring := lengthOfLongestSubstring("abcabcbb")
	if substring != 3 {
		t.Error("abcabcbb should 3, but:", substring)
	}

	substring = lengthOfLongestSubstring("bbbbb")
	if substring != 1 {
		t.Fatal("bbbbb should 1, but:", substring)
	}

	substring = lengthOfLongestSubstring("pwwkew")
	if substring != 3 {
		t.Fatal("pwwkew should 3, but:", substring)
	}

	substring = lengthOfLongestSubstring("eeydgwdykpv")
	if substring != 7 {
		t.Fatal("eeydgwdykpv should 7, but:", substring)
	}
}
