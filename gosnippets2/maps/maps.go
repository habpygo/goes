package main

import "fmt"

var myHashTable = make(map[string]int)

func main() {
	fmt.Println("This is just a test")
	myHashTable["Howdy"] = 34
	myHashTable["John"] = 55

	addToMap("gosh", 44)
	removeFromMap("John")

	for k, v := range myHashTable {
		fmt.Println(k)
		fmt.Println(v)
	}
}

func addToMap(a string, b int) {
	myHashTable[a] = b
}

func removeFromMap(c string) {
	delete(myHashTable, c)
}
