package main

import (
	"fmt"
)

type myStructure struct {
	Name    string
	Surname string
	Height  int32
}

func createStruct(n, s string, h int32) *myStructure {
	if h > 300 {
		h = 0
	}
	return &myStructure{n, s, h}
}

func retStructure(n, s string, h int32) myStructure {
	if h > 300 {
		h = 0
	}
	return myStructure{n, s, h}
}

func main() {
	s1 := createStruct("Mihalis", "Tsoukalos", 123)
	s2 := retStructure("Mihalis", "Tsoukalos", 123)

	fmt.Println((*s1).Name)
	// Mihalis
	fmt.Println(s2.Name)
	// Mihalis
	fmt.Println(s1)
	// &{Mihalis Tsokalos 123}
	fmt.Println(s2)
	// {Mihalis Tsokalos 123}
}
