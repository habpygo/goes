package main

import (
	"fmt"
)

var a string
var b string

func main() {

	a = "hello"
	b = a
	a = "twello"

	fmt.Println("a is: ", b)
	fmt.Println(string(65))     // "A", not "65"
	fmt.Println(string(0x4eac)) //

}
