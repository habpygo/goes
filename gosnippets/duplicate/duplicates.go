/*
Design an algorithm and write code to remove the duplicate characters in a string
without using any additional buffer. NOTE: One or two additional variables are fine.
An extra copy of the array is not
*/

package main

import (
	"fmt"
	"sort"
)

func removeDuplicates(input []string) []string {
	mapValues := map[string]bool{}
	result := []string{}

	for v := range input {
		mapValues[input[v]] = true
	}
	for key, _ := range mapValues { // key is the string!
		result = append(result, key)
	}

	return result
}

// removes duplicates by first sorting the string slice and then removing one of the doubles
// a double exists when there are two consecutive strings that are equal
// by appending only the strings that are not double
// the last string has to be added
// Note: original order is not maintained
func removeSortDuplicates(input []string) []string {
	newStringSlice := []string{}
	sort.Strings(input) // []string input is now sorted and we can compare

	for i := 0; i < len(input)-1; i++ {
		if input[i] != input[i+1] {
			newStringSlice = append(newStringSlice, input[i])
		}
	}
	newStringSlice = append(newStringSlice, input[len(input)-1])

	return newStringSlice
}

func main() {

	rawstring := []string{"Hi there", "How are you?", "Is this real?", "I don't know", "Is this real?", "How are you!", "Hi there"}
	rawstring2 := []string{"Hi", "there", "is", "is", "this", "there", "where", "real", "world", "world", "says", "Hi", "there"}

	fmt.Println("From: removeDuplicates: nieuw string is: ", removeDuplicates(rawstring2))
	fmt.Println("From: removeDuplicates: nieuw string is: ", removeDuplicates(rawstring))

}
