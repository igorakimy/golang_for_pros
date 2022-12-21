package main

import (
	"fmt"
)

func main() {
	aMap := make(map[string]int)
	// var aMap map[string]int
	aMap = nil
	fmt.Println(aMap)
	aMap["test"] = 1
	// map[]
	// panic: assignment to entry in nil map
}
