package main

import "fmt"

func main() {

	myS := "Hi there how is it going?"
	mySliceS := CreateStringQueue(myS) // returns []byte
	fmt.Println("mySliceS is: ", string(mySliceS))

	addS := "What in the world "
	enq := Enqueue(mySliceS, addS)
	fmt.Println("Enqueued is: ", string(enq))

	byteEnq := []byte(enq)
	deq := Dequeue(byteEnq, 5)
	fmt.Println("Dequeued is: ", string(deq))

}

func CreateStringQueue(s string) []byte {
	byteQuey := []byte(s)
	return byteQuey
}

func Enqueue(q []byte, s string) string {
	// length of queue must be extended to cater for the new string
	totLength := len(q) + len(s)
	newQueue := make([]byte, totLength, 2*totLength+1) // create some extra capacity
	newS := []byte(s)
	copy(newQueue[:len(s)], newS)
	copy(newQueue[len(s):], q)

	return string(newQueue)
}

func Dequeue(s []byte, noOfChars int) string {

	ds := s[:len(s)-noOfChars]
	return string(ds)
}

/*
○ → go run queue.go
mySliceS is:  Hi there how is it going?
Enqueued is:  What in the world Hi there how is it going?
Dequeued is:  What in the world Hi there how is it g
*/
