package main

import (
	"fmt"
)

func main() {
	aSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(aSlice)
	// [1 2 3 4 5]
	integers := make([]int, 2)
	fmt.Println(integers)
	// [0 0]
	integers = nil
	fmt.Println(integers)
	// []

	anArray := [5]int{-1, -2, -3, -4, -5}
	refAnArray := anArray[:]
	fmt.Println(anArray)
	// [-1 -2 -3 -4 -5]
	fmt.Println(refAnArray)
	// [-1 -2 -3 -4 -5]
	anArray[4] = -100
	fmt.Println(refAnArray)
	// [-1 -2 -3 -4 -100]

	s := make([]byte, 5)
	fmt.Println(s)
	// [0 0 0 0 0]
	twoD := make([][]int, 3)
	fmt.Println(twoD)
	// [[] [] []]
	fmt.Println()

	// Инициализация всех элементов двумерного среза вручную
	for i := 0; i < len(twoD); i++ {
		for j := 0; j < 2; j++ {
			twoD[i] = append(twoD[i], i*j)
		}
	}

	fmt.Println(twoD)

	for _, x := range twoD {
		for i, y := range x {
			fmt.Println(i, y)
		}
		fmt.Println()
	}
}
