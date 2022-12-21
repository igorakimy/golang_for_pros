package main

import (
	"fmt"
	"unicode"
)

func main() {
	const sL = "\x99\x00ab\x50\x00\x23\x50\x29\x9c"

	for i := 0; i < len(sL); i++ {
		// Проверить, можно ли напечатать руну
		if unicode.IsPrint(rune(sL[i])) {
			fmt.Printf("%c\n", sL[i])
		} else {
			fmt.Println("Not printable!")
		}
	}
	// Not printable!
	// Not printable!
	// a
	// b
	// P
	// Not printable!
	// #
	// P
	// )
	//Not printable!
}
