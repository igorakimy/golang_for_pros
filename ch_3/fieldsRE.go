package main

import (
	"fmt"
	"os"
	"strings"
)

func matchRecord(s string) bool {
	fields := strings.Split(s, ",")
	if len(fields) != 3 {
		return false
	}

	if !matchNameSur(fields[0]) {
		return false
	}

	if !matchNameSur(fields[1]) {
		return false
	}

	return matchInt(fields[2])
}

func main() {
	args := os.Args

	if len(args) == 1 {
		fmt.Println("usage: fieldsRE name,sure,phone")
		os.Exit(1)
	}

	record := args[1]

	fmt.Println(matchRecord(record))
}
