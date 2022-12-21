package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3}
	a := [3]int{4, 5, 6}

	ref := a[:]
	fmt.Println("Existing array:\t", ref)
	// [4 5 6]
	t := append(s, ref...)
	fmt.Println("New slice:\t", t)
	// [1 2 3 4 5 6]
	s = append(s, ref...)
	fmt.Println("Existing slice:\t", s)
	// [1 2 3 4 5 6]
	s = append(s, s...)
	fmt.Println("s+s:\t\t", s)
	// [1 2 3 4 5 6 1 2 3 4 5 6]
}
