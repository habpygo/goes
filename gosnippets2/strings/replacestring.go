package main

import (
	"fmt"
)

func main() {

	fmt.Println(replaceString("Hallo hoe gaat het met jou beste man?"))
}

func replaceString(s string) string {
	r := []byte("%20")
	b := make([]byte, 2*len(s))
	b = []byte(s)
	c := make([]byte, 2*len(s))

	for i := 0; i < len(b); i = i + 1 {
		if string(b[i]) == " " {
			copy(c[i:], r)
			copy(c[i+3:], b[i+1:]) // % 2 0 are three characters
			copy(b[0:], c[0:])

		} else {
			c[i] = b[i] // i's should be different
		}

	}
	return string(c)
}
