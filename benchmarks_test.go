package strutil

import (
	"testing"
)

func BenchmarkPalindromeLarge(b *testing.B) {

	val := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	b.SetBytes(int64(len([]byte(val))))
	for n := 0; n < b.N; n++ {
		IsPalindrome(val)
	}
}

func BenchmarkPalindromeMid(b *testing.B) {

	val := "tattarrattat"

	b.SetBytes(int64(len([]byte(val))))
	for n := 0; n < b.N; n++ {
		IsPalindrome(val)
	}
}

func BenchmarkPalindromeSmall(b *testing.B) {

	val := "racecar"

	b.SetBytes(int64(len([]byte(val))))
	for n := 0; n < b.N; n++ {
		IsPalindrome(val)
	}
}

func BenchmarkPalindromeMultibyte(b *testing.B) {

	val := "♥☺♥☺♥☺♥"

	b.SetBytes(int64(len([]byte(val))))
	for n := 0; n < b.N; n++ {
		IsPalindrome(val)
	}
}

func BenchmarkAnagramLarge(b *testing.B) {

	val1 := "Eleven plus two"
	val2 := "Twelve plus one"

	b.SetBytes(int64(len([]byte(val1)) + len([]byte(val2))))
	for n := 0; n < b.N; n++ {
		IsAnagram(val1, val2)
	}
}

func BenchmarkAnagramMid(b *testing.B) {

	val1 := "A gentleman"
	val2 := "Elegant man"

	b.SetBytes(int64(len([]byte(val1)) + len([]byte(val2))))
	for n := 0; n < b.N; n++ {
		IsAnagram(val1, val2)
	}
}

func BenchmarkAnagramSmall(b *testing.B) {

	val1 := "Race car"
	val2 := "Car race"

	b.SetBytes(int64(len([]byte(val1)) + len([]byte(val2))))
	for n := 0; n < b.N; n++ {
		IsAnagram(val1, val2)
	}
}

func BenchmarkAnagramMultibyte(b *testing.B) {

	val1 := "♥☺♥☺♥☺♥"
	val2 := "♥♥♥♥☺☺☺"

	b.SetBytes(int64(len([]byte(val1)) + len([]byte(val2))))
	for n := 0; n < b.N; n++ {
		IsAnagram(val1, val2)
	}
}
