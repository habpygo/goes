package testfolder

import (
	"testing"
)

func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] { // as always: [0,1,2,3,4] len is 5, laatste is len[s]-1
			return false
		}
	}
	return true
}

func TestPalindrome(t *testing.T) {
	if !IsPalindrome("detartrated") {
		t.Error(`IsPalindrome("detartrated") = false`)
	}
	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome("kayak") = false`)
	}
}

func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("palindrome") {
		t.Error(`IsPalindrome("palindrome") = true`)
	}
}
