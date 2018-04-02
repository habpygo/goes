package main

import "fmt"

var n int = 0
var lowest = 0
var highest = 1000

func BinarySearch(values []int, value int) {

	middleValue := (lowest + highest) / 2

	for _, v := range values {
		if v < value {
			lowest = v
		}
		if v > value {
			highest = v
		}
		if v == value {
			fmt.Println("Success! The found number is: ", v)

		}
		middleValue = (lowest + highest) / 2
	}
}

func main() {
	searchValues := []int{3, 4, 6, 7, 9, 44, 5, 10, 12}
	searchValue := 9
	BinarySearch(searchValues, searchValue)

}
