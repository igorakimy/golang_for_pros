package main

import (
	"fmt"
)

func main() {
	s1 := make([]int, 5)
	reSlice := s1[1:3]
	fmt.Println(s1)
	// [0 0 0 0 0]
	fmt.Println(reSlice)
	// [0 0]
	reSlice[0] = -100
	reSlice[1] = 123456
	fmt.Println(s1)
	// [0 -100 123456 0 0]
	fmt.Println(reSlice)
	// [-100 123456]
}
