package strutil

import "fmt"

func ExampleIsPalindrome() {

	res := IsPalindrome("racecar")

	fmt.Println(res)
	// Output: true
}

func ExampleIsAnagram() {

	res := IsAnagram("Madam Curie", "Radium came")

	fmt.Println(res)
	// Output: true
}
