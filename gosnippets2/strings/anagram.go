package main

import (
	"fmt"
)

func main() {

	var var1 = "Hallo hoe gaat het"
	var var2 = "HaelleaH"

	fmt.Println(checkAnagram(var1))
	fmt.Println(checkAnagram(var2))
}

func checkAnagram(s string) bool {
	for i, j := 0, len(s)-1; i < len(s); i, j = i+1, j-1 {
		if s[i] == s[j] {
			return true
		}
	}
	return false
}
