package main

import "fmt"

var newArray [10]int // or
var newArray2 = [10]int{}
var stringSlice []string

func main() {
	s := make([]string, 10, 10)
	s[0] = "hallo"
	s[1] = "dumbass"
	s[2] = "you"
	s[3] = "there"
	s[4] = "from beyond"
	s[5] = "does it"
	s = append(s, "me", "what", "now", "happens", "tomorrow")
	s = append(s, "me", "what", "now", "happens", "tomorrow")
	s = append(s, "me", "what", "now", "happens", "tomorrow")
	s = append(s, "so what now")
	fmt.Println(s)
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)

	// This will frop the first and last element of the slice = pops and deques
	c = c[1 : len(c)-1]
	fmt.Println(c)

}

func createStringSlice() {
	//Stringslice := make([]string, 0, 10)
}

func addToStringSlice(slice []string, item ...string) {

}

func AppendInt(slice []int, data ...int) []int {
	m := len(slice)     // get the length of the existing slice
	n := m + len(data)  // get the length of the existing slice + the newly to be added data slice
	if n > cap(slice) { // if n exeeds capacity =
		newSlice := make([]int, (n+1)*2) // create a slice with twice the capacity with some extra
		copy(newSlice, slice)            // copy the slice into the newly create slice
		slice = newSlice                 // and set the existing slice to the new slice
	}
	slice = slice[0:n]     // create the slice with the new length
	copy(slice[m:n], data) // and copy the data into the slice from the old position m till the new position n
	return slice
}
