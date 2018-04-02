package main

import "fmt"

var x = 10
var y = 5
var n = 12

func main() {
	f := fibonacci()
	for i := 0; i < n; i++ {
		fmt.Println(f())
	}
	fib(n)
}

func fibonacci() func() int {
	x, y = 1, 0 // usually it's x, y = 0, 1  but this way we can get the 0 value as well
	return func() int {
		x, y = y, x+y
		return x
	}
}

// computing the n-th Fibonacci number iteratively:
func fib(n int) int {
	x, y := 1, 0 // usually it's x, y = 0, 1  but this way we can get the 0 value as well
	for i := 0; i < n; i++ {
		x, y = y, x+y
		fmt.Println("The nth fibonacci number is is: ", x, "y is: ", y)
	}
	return x
}
