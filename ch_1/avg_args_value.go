package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var sum float64
	var floatArgsCount int
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please give one or more floats")
		os.Exit(1)
	}
	for i := 1; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err != nil {
			continue
		}
		sum += n
		floatArgsCount++
	}

	avg := sum / float64(floatArgsCount)
	fmt.Println("Avg:", avg)
}
