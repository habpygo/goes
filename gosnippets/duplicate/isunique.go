package main

import "fmt"

// Function isUnique returns a boolean value indicating if a string has all unique characters

func IsUnique(s string) (int, bool) {
	m := make(map[rune]bool, len(s))
	var index int
	for index, r := range s {
		if m[r] {
			return index, false
		}
		m[r] = true
	}
	return index, true
}

func main() {
	s := "Hey JoHn"
	index, value := IsUnique(s)
	if value {
		fmt.Println("All characters are unique!")
	} else {
		fmt.Printf("Found non-unique character at index %v in string %v", index, s)
	}
}
