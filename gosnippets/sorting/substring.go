package main

import (
	"fmt"
	"sort"
	"strings"
)

func isSubstring(inputstring, substring string) bool {

	if strings.Contains(inputstring, substring) {
		return true
	}
	return false
}

//SORTING STRING FIRST, THEN COMPARE

func sortString(s string) string {
	w := strings.Split(s, "")
	sort.Strings(w)
	return strings.Join(w, "")
}

func main() {
	s1 := "droplol"
	s2 := "lropdul"
	sortedString1 := sortString(s1)
	sortedString2 := sortString(s2)

	if sortedString1 == sortedString2 {
		fmt.Println("The strings: ", s1, "and: ", s2, "are anagrams")
	} else {
		fmt.Println("The strings are no anagram")
	}
	fmt.Println(isSubstring(s1, s2))
}
