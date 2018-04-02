package main

import (
	"fmt"
)

func main() {
	myMap := make(map[int]string)

	myMap[1] = "man"
	myMap[2] = "woman"
	myMap[3] = "dog"
	myMap[4] = "Mechelaar"
	myMap[5] = "dog"

	fmt.Println("Number 4 is a: ", myMap[4])
	for k, v := range myMap {
		fmt.Println("Voor de key: ", k, "krijgen we de waarde: ", v)
		fmt.Println(len(myMap) == 0)

	}

}
