package main

import (
	"fmt"
	"os"
	"regexp"
)

func matchNameSur(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[A-Z][a-z]*$`)
	return re.Match(t)
}

func main() {
	args := os.Args

	if len(args) == 1 {
		fmt.Println("usage: numSureRE {name}")
		os.Exit(1)
	}

	nameSur := args[1]

	fmt.Println(matchNameSur(nameSur))
}
