package main

import (
	"fmt"
)

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func rotateSlice(s []int) []int {
	fmt.Println(s)
	reverse(s[:2]) // [2 1 3 4 5 6 7 8 9]
	fmt.Println(s)
	reverse(s[2:]) // [2 1 9 8 7 6 5 4 3]
	fmt.Println(s)
	reverse(s) // 3 4 5 6 7 8 9 1 2]
	fmt.Println(s)

	return s
}

// For equality tests do the following
/*
Slice values are deeply equal when all of the following are true:
1. they are both nil or both non-nil,
2. they have the same length, and either
   2a. they point to the same initial entry of the same underlying array (that is, &x[0] == &y[0]) or
   2b. their corresponding elements (up to length) are deeply equal.
*/
func equal(x, y []int) bool {
	if len(x) != len(y) {
		return false // if the length is not the same we can stop right there
	}
	// for i := range x {
	// 	if x[i] != y[i] {
	// 		return false
	// 	}
	// }
	for i, v := range x {
		if v != y[i] {
			return false
		}
	}
	return true
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice2 := rotateSlice(slice)
	fmt.Println(slice2)
	//slice3 := []int{3, 4, 5, 6, 7, 8, 9, 1, 2}
	fmt.Println("Slices are equal? ", equal(slice, slice2))
}
