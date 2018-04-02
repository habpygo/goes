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

}
