package main

import (
	"fmt"
	"sort"
)

func main() {
	// s := make(map[int]bool) but we prefer here to use a map literal
	s := map[int]bool{5: true, 3: true, 10: true, 12: true, 14: true}
	s1 := map[int]bool{1: true, 2: true, 4: true, 10: true, 12: true, 34: true, 45: true, 67: true, 57: true}
	Unite(s, s1)
	fmt.Println("Intersectin is: ", Intersection(s, s1))
	fmt.Println("Difference is: ", Difference(s, s1))
}

// Unite two sets thereby deleting doubles
func Unite(a, b map[int]bool) {
	s_Union := map[int]bool{}
	for k := range a {
		s_Union[k] = true
	}
	for k, _ := range b {
		s_Union[k] = true
	}
	for x, y := range s_Union {
		fmt.Println(x, y)
	}
	sortedInts := SortSet(s_Union)
	fmt.Println(sortedInts)
}

// Intersection - Given two sets, this function returns another set that has all items that are part of both sets.
func Intersection(m, n map[int]bool) map[int]bool {
	section := map[int]bool{}
	for k, _ := range m {
		for l, _ := range n {
			if k == l {
				section[k] = true
			}
		}
	}
	return section
}

// s: 5 3 10 12 14
// s1: 1 2 4 10 12 34 45 67 57  //antwoord: 5 3 14
// Difference - This returns a list of items that are in one set but NOT in a different set.
func Difference(m, n map[int]bool) map[int]bool {
	difference := map[int]bool{}

	for k := range m {
		for l := range n {
			if k != l {
				difference[k] = true
			} 

			}

		}
	}

	return difference
}

// Subset - returns a boolean value that shows if all the elements in one set are included in a different set.
func Subset(m map[int]bool, n map[int]bool) bool {
	return false
}

func SortSet(m map[int]bool) []int {

	intSlice := []int{}
	for k := range m {
		intSlice = AppendInt(intSlice, k)
	}

	sort.Ints(intSlice)
	return intSlice
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
