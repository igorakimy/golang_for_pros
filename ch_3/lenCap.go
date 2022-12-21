package main

import (
	"fmt"
)

func printSlice(x []int) {
	for _, number := range x {
		fmt.Print(number, " ")
	}
	fmt.Println()
}

func main() {
	aSlice := []int{-1, 0, 4}
	fmt.Printf("aSlice: ")
	printSlice(aSlice)
	// aSlice: -1 0 4
	fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice))
	// Cap: 3, Length: 3
	aSlice = append(aSlice, -100)
	fmt.Printf("aSlice: ")
	printSlice(aSlice)
	// aSlice: -1 0 4 -100
	fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice))
	// Cap: 6, Length: 4

	aSlice = append(aSlice, -2)
	aSlice = append(aSlice, -3)
	aSlice = append(aSlice, -4)
	printSlice(aSlice)
	// -1 0 4 -100 -2 -3 -4
	fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice))
	// Cap: 12, Length: 7
}
