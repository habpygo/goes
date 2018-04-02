package main

import (
	"fmt"
)

//const reversestring = "We are going to learn a lot from this exercise!"
const reversestring = "!esicrexe siht morf tol a nrael ot gniog era eW"

func main() {

	fmt.Println(ReverseString(reversestring))
	fmt.Println(ReverseRune(reversestring))
}

func ReverseRune(s string) string {
	n := 0
	rune := make([]rune, len(s)) // why make this rune
	for _, r := range s {
		rune[n] = r
		n++
	}
	//rune = rune[0:n] // slice all other characters away - not really necessary
	// Now reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	return string(rune)
}

func ReverseString(s string) []string {
	n := 0
	rs := make([]string, len(s)) // why make this rune
	for _, r := range s {
		rs[n] = string(r)
		n++
	}
	rs = rs[0:n] // slice all other characters away - not really necessary
	// Now reverse
	for i := 0; i < n/2; i++ {
		rs[i], rs[n-1-i] = rs[n-1-i], rs[i]
	}
	return rs
}
