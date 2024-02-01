package main

import (
	"fmt"
	"os"
	"regexp"
)

func matchInt(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[-+]?\d+$`)
	return re.Match(t)
}

func main() {
	args := os.Args

	if len(args) == 1 {
		fmt.Println("usage: intRE {number}")
		os.Exit(1)
	}

	num := args[1]

	fmt.Println(matchInt(num))
}
