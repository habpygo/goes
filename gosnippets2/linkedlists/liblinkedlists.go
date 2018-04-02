package main

import (
	"container/list"
	"fmt"
)

func main() {
	// Create a new list and put some numbers in it.
	/*
		3  == Insert before eHead
		0  == PushFront (eHead)
		5  == PushBack (eLast)
		25 == after eLast after 2 has been inserted as after eLast
		2  == after elast and becomes eLast

	*/
	l := list.New()
	eLast := l.PushBack(5)
	eHead := l.PushFront(0)
	l.InsertAfter(2, eLast)
	l.InsertBefore(3, eHead)
	l.InsertAfter(25, eLast)
	eMiddle := l.InsertAfter(20, eHead)
	l.InsertBefore(23, eMiddle)
	l.InsertAfter(21, eMiddle)
	l.InsertAfter(2, eMiddle)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("->%+v", e.Value)
	}

}
