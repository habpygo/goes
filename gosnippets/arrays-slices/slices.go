package main

import (
	"fmt"
)

var d = []int{1, 2, 3, 4, 5}

func makeMyIntSlice(elements []int) []int {
	s := make([]int, 10, 10)
	s = append(s, elements...)
	fmt.Println("totale lengte is: ", len(s)+len(d))

	return s
}

func main() {

	mySlice := makeMyIntSlice(d)
	fmt.Println("mySlice is: ", mySlice)

}
