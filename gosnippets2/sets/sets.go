// model a set from a map
package main

import "fmt"

func main() {
	set := NewIntSet()
	set.Add(1)
	set.Add(17)
	set.Add(23)
	fmt.Println("size of IntSet is: ", set.Size())
	fmt.Println("set contains 17? ", set.Contains(17))
	set.RemoveSet(17)
	fmt.Println("size of IntSet is: ", set.Size())
	fmt.Println("set contains 17? ", set.Contains(17))

}

type IntSet struct {
	set map[int]bool
}

func NewIntSet() *IntSet {
	return &IntSet{make(map[int]bool)}
}

func (set *IntSet) Add(i int) bool {
	v, found := set.set[i] // found gets true or false
	set.set[i] = true
	fmt.Println("v is: ", v)
	return !found // false if it exists
}

func (set *IntSet) Contains(i int) bool {
	_, found := set.set[i]
	return found // true if it existed already
}

func (set *IntSet) RemoveSet(i int) {
	_, found := set.set[i]
	if found {
		delete(set.set, i)
	}
}

func (set *IntSet) Size() int {
	return len(set.set)
}
