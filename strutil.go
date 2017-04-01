package strutil

import (
	"math"
	"strings"
	"unicode/utf8"
)

// IsPalindrome determines if the provided string is a palindrome.
func IsPalindrome(s string) bool {

	// after some testing this seems to be tuned properly.
	if len(s) > 7 && len(s) < 33 {
		return isPalindromeSmall(s)
	}

	return isPalindromeMid(s)
}

// isPalindromeSmall determines if the provided string is a palindrome for a small string.
func isPalindromeSmall(s string) bool {

	runes := []rune(s)
	j := len(runes) - 1

	for i := 0; i < len(runes)/2; i++ {

		if runes[i] != runes[j] {
			return false
		}
		j--
	}

	return true
}

// isPalindromeMid determines if the provided string is a palindrome for a mid-large string.
func isPalindromeMid(s string) bool {

	var r1, r2 rune
	var size int

	for len(s) > 1 { // > 1 to handle odd lengths

		r1, size = utf8.DecodeRuneInString(s)
		s = s[size:]

		r2, size = utf8.DecodeLastRuneInString(s)

		if r1 != r2 {
			return false
		}

		s = s[:len(s)-size]
	}

	return true
}

// IsAnagram determines if the provided string is an anagram of the supplied subject.
func IsAnagram(subject string, s string) bool {

	if len(subject) != len(s) || subject == s {
		return false
	}

	subject = strings.ToLower(subject)
	s = strings.ToLower(s)

	m := make(map[rune]uint64, int(math.Min(float64(len(subject)), 64)))

	for _, r := range subject {

		val, found := m[r]
		if found {
			val++
			m[r] = val
			continue
		}

		m[r] = 1
	}

	for _, r := range s {

		val, found := m[r]
		if !found {
			return false
		}

		if val == 1 {
			delete(m, r)
			continue
		}

		val--
		m[r] = val
	}

	return len(m) == 0
}
