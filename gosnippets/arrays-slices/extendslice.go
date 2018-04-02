package main

import "fmt"

/*
Eerst wordt de capaciteit gecontroleerd. Als hij vol is, breiden we hem
2 keer uit. Dan kopieren we de originele slice in de nieuwe zetten de oude
slice weer terug. Dan voegen we een 0 toe aan de slice en vervangen die door
de nieuwe waarde.
*/

var valueToExtent int = 123456789
var sliceToAppend = []int{12, 13, 14, 15, 16}

func main() {
	slice := []int{4, 5, 4, 2, 78, 234}
	slice = Extent(slice, valueToExtent)
	fmt.Println("extended slice is: ", slice)
	slice = Append(slice, sliceToAppend...)
	fmt.Println("appended slice is: ", slice)

}

func Extent(s []int, v int) []int {
	n := len(s)
	if n == cap(s) { // slice is full
		newSlice := make([]int, len(s), 2*len(s)+1)
		copy(newSlice, s)
		s = newSlice
	}
	s = s[0 : n+1] // append a 0 to the existing slice
	fmt.Println("s is: ", s)
	s[n] = v // now replace the 0 with the to be added value
	return s
}

func Append(s []int, elements ...int) []int {
	n := len(s)
	neededLength := len(s) + len(elements)
	if neededLength > cap(s) {
		newSlice := make([]int, neededLength, 2*neededLength+1)
		copy(newSlice, s)
		s = newSlice
	}
	s = s[:neededLength]
	copy(s[n:], elements)

	return s
}
