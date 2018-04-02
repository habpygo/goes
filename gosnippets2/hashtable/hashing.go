package main

import (
	"fmt"
)

func main() {
	fmt.Println("The Hash is: ", hash("Harry Boer"))
	fmt.Println("The Hash is: ", hash("Henri Boer"))
}

// the hash() private function uses the famous Horner's method
// to generate a hash of a string with O(n) complexity
func hash(k string) int {
	key := fmt.Sprintf("%s", k)
	h := 0
	for i := 0; i < len(key); i++ {
		h = 31*h + int(key[i]) // int(key[i]) gives the ASCII number for the character
	}
	return h
}

/*
prints:
The Hash is:  1324849107223644594
The Hash is:  1325957397764573618
*/
