package main

import "fmt"

func main() {

	x := 120
	y := 100

	g := gdc(x, y)

	fmt.Printf("gdc of %v and %v is %v", x, y, g)

}

func gdc(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
