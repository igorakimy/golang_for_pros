package main

import (
	"fmt"
	"os"
	"regexp"
)

func findWrongPartsIP(input string) []string {
	partIP := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	grammar := partIP + "\\." + partIP + "\\." + partIP + "\\." + partIP
	re := regexp.MustCompile(grammar)
	return re.FindStringSubmatch(input)
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an IP address!")
		return
	}

	ip := findWrongPartsIP(arguments[1])
	fmt.Println(ip)
}
