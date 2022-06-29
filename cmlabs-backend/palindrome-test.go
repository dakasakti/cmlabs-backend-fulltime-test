package main

import (
	"fmt"
)

func palindrome(word string) bool {
	for i := 0; i < len(word)/2; i++ {
		if word[i] != word[len(word)-1-i] {
			return false
		}
	}

	return true
}

func main() {
	// Single
	fmt.Println(palindrome("SAIPPUAKIVIKAUPPIAS"))
	fmt.Println(palindrome("Aibohphobia"))
	fmt.Println(palindrome("Anna"))
	fmt.Println(palindrome("Civic"))

	// Multiple
	fmt.Println(palindrome("My gym"))
	fmt.Println(palindrome("No lemon, no melon"))
}
