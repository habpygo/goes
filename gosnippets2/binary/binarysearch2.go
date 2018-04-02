package main

import "errors"

func main() {
	searchArray := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91} // note that array is sorted from low to high
	searchItem := 56

}

type Node struct {
	Value string
	Data  string
	Left  *Node
	Right *Node
}

func (n *Node) Insert(value, data string) error {
	if n == nil {
		return errors.New("Cannot insert a value into a nil tree")
	}

	switch {
	case value == n.Value: // If the data is already in the tree, return.
		return nil

	case value < n.Value: // If the data value is less than the current node’s value,
		if n.Left == nil { // and if the left child node is nil,
			n.Left = &Node{Value: value, Data: data} // fill left child node.
			return nil
		}
		return n.Left.Insert(value, data) // Else call Insert on the left subtree.

	}
}
