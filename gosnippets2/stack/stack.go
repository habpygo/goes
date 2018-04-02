package main

import "fmt"

func main() {
	fmt.Println("stack is: ", createStack())
	myStack := createStack()
	fmt.Println("pushed on stack: ", pushOnStack(345, myStack))
	fmt.Println("popped from stack: ", popFromStack(myStack))

}

func createStack() []int {
	stack := []int{3, 4, 5, 6, 7, 8, 6, 9, 12}
	return stack
}

func pushOnStack(i int, st []int) []int {
	stack := st
	stack = append(st, i)

	return stack
}

func popFromStack(st []int) []int {
	stack := st
	stack = stack[0:len(stack)]
	return stack
}
