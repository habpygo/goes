// Go's `sort` package implements sorting for builtins
// and user-defined types. We'll look at sorting for
// builtins first.

package main

import (
	"fmt"
	"math/rand"
)

func bubbleIntSort(arrayToSort []int) {
	bubbling := true
	for bubbling { // bubbling is true, dus we gaan de for loop in
		bubbling = false                          // maar nu is ie false
		for i := 0; i < len(arrayToSort)-1; i++ { // loopt hele array af en kijkt of er twee te swappen zijn
			if arrayToSort[i] > arrayToSort[i+1] { // zo niet, dan is de array al gesorteerd, zo ja, dan wordt
				swapints(arrayToSort, i, i+1) // bubbling true gezet en kan de for loop weer opnieuw beginnen
				bubbling = true
			}
		}
	}
}

func bubbleCharSort(charsToSort []byte) {
	bubbling := true
	for bubbling {
		bubbling = false
		for i := 0; i < len(charsToSort)-1; i++ {
			if charsToSort[i] > charsToSort[i+1] {
				swapchars(charsToSort, i, i+1)
				bubbling = true
			}
		}
	}
}

func swapchars(input []byte, i, j int) {
	tmp := input[i]
	input[i] = input[j]
	input[j] = tmp
}

func swapints(input []int, i, j int) {
	temp := input[i]
	input[i] = input[j]
	input[j] = temp
}

func main() {
	// rand.Perm returns, as a slice of n ints, a pseudo-random permutation of the integers [0,n) from the default Source.
	intArraySort := make([]int, 20)
	intArraySort = rand.Perm(20)
	fmt.Println("The original unsorted integer array is: ", intArraySort)
	bubbleIntSort(intArraySort)
	fmt.Println("The sorted integer array is: ", intArraySort)
	fmt.Println("-------------------")
	// strings
	myString := "Hallo wat is dit leuk!"
	stringArraySort := []byte(myString)
	fmt.Println("characters to sort is: ", string(stringArraySort))
	fmt.Println("byte slice of the characters is: ", stringArraySort)
	fmt.Println("de 5e character in de string is myString[6]", string(myString[6]))
	bubbleCharSort(stringArraySort)
	fmt.Println("The sorted character array is: ", string(stringArraySort))

}
