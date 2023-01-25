package main

import (
	"fmt"
)

func add(x, y int64) int64 {
	return x + y
}

func main() {
	fmt.Println(add(1, 2))
}
