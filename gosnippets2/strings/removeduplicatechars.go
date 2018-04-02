package main

import (
	"fmt"
)

func main() {

	fmt.Println(removeDupChars("AlhoewelAbedoelah"))
}

func removeDupChars(s string) string { // s t r r i n g
	r := []rune(s)                // first make rune
	for i := 0; i < len(r); i++ { // note: a) i < len(s) of i = len(s)-1
		for j := i + 1; j < len(r); j++ {
			if r[j] == r[i] {
				r = append(r[0:j], r[j+1:]...)
			}
		}
	}
	return string(r) // back to string
}
