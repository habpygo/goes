package main

import "fmt"

func main() {
	fmt.Println("result is: ", callFactorial(5))
	testArray := []int{3, 4, 5, 6, 7, 8, 10}
	fmt.Println("callSum is: ", callSum(testArray))
	fmt.Println("recursionSum is: ", recursionSum(testArray))

}

// example of recursion
func callFactorial(x int) int {
	var result int
	if x == 1 {
		return 1
	}
	result = x * callFactorial(x-1) // stays here and puts new result on the stack
	return result
}

// with recursion
func recursionSum(x []int) int {
	//var result int

	if len(x) == 1 {
		return x[0]
	}
	return x[len(x)-1] + recursionSum(x[:len(x)-1])
}

// without recursion
func callSum(x []int) int {
	var result int
	fmt.Println(x)
	for k, v := range x {
		fmt.Println(k, v)
		result = result + v
	}
	return result
}

/*
○ → go run factorial.go
result is:  120
[3 4 5 6 7 8 10]
0 3
1 4
2 5
3 6
4 7
5 8
6 10
callSum is:  43
recursionSum is:  43
*/
