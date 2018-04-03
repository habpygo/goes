package main

import (
	"bytes"
	"fmt"
)

// intsToString is like fmt.Sprintf(values) but adds commas.
func intsToString(values []int) string {
	var buf bytes.Buffer

	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

// recursive comma function
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

// comma function without recursion
func commaWithout(s string) string {
	var buf bytes.Buffer
	for k, v := range s {
		if (k+2)%3 == 0 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%s", string(v))
	}
	return buf.String()
}

func main() {
	a := [3]int{1, 2, 3} // array literal
	for k, v := range a {
		fmt.Printf("%d %d\n", k, v)
	}
	fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"
	fmt.Println("output from comma is: ", comma("Hello there how is it?"))
	fmt.Println("output from commaWithout is: ", commaWithout("Hello there how is it?"))
	fmt.Println("output from commaWithout is: ", commaWithout("Hello there how is 345,00 it?"))
}
