package main

import (
	"fmt"
)

const input = "We are going to learn a lot from this exercise."

func main() {

	fmt.Println("IsUniqueChar: ", string(UniqueBytes(input)))
}

func UniqueBytes(s string) []byte {
	occur := [256]byte{}
	order := make([]byte, 0, 256)
	n := 0
	for i := 0; i < len(s); i++ {
		b := s[i]
		switch occur[b] {
		case 0:
			occur[b] = 1
			order = append(order, b)
			n++
		case 1:
			occur[b] = 2
			n--
		}
	}
	if n == 0 {
		return nil
	}
	result := make([]byte, 0, n)
	for _, b := range order {
		if occur[b] == 1 {
			result = append(result, b)
		}
	}
	return result
}
