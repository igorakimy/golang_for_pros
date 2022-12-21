package main

import (
	"fmt"
)

func main() {
	const sLiteral = "\x99\x42\x32\x55\x50\x35\x23\x50\x29\x9c"
	fmt.Println(sLiteral)
	// �B2UP5#P)�
	fmt.Printf("x: %x\n", sLiteral)
	// x: 9942325550352350299c
	fmt.Printf("sLiteral length: %d\n", len(sLiteral))
	// sLiteral length: 10

	for i := 0; i < len(sLiteral); i++ {
		fmt.Printf("%x ", sLiteral[i])
	}
	// 99 42 32 55 50 35 23 50 29 9c
	fmt.Println()
	fmt.Printf("q: %q\n", sLiteral)
	// q: "\x99B2UP5#P)\x9c"
	fmt.Printf("+q: %+q\n", sLiteral)
	// +q: "\x99B2UP5#P)\x9c"
	fmt.Printf("x: % x\n", sLiteral)
	// x: 99 42 32 55 50 35 23 50 29 9c
	fmt.Printf("s: As a string: %s\n", sLiteral)
	// s: As a string: �B2UP5#P)�

	s2 := "€£³"
	for x, y := range s2 {
		fmt.Printf("%#U starts at byte position %d\n", y, x)
	}
	// U+20AC '€' starts at byte position 0
	// U+00A3 '£' starts at byte position 3
	// U+00B3 '³' starts at byte position 5
	fmt.Printf("s2 length: %d\n", len(s2))
	// s2 length: 7

	const s3 = "ab12AB"
	fmt.Println("s3:", s3)
	// s3: ab12AB
	fmt.Printf("x: % x\n", s3)
	// x: 61 62 31 32 41 42
	fmt.Printf("s3 length: %d\n", len(s3))
	// s3 length: 6
	for i := 0; i < len(s3); i++ {
		fmt.Printf("%x ", s3[i])
	}
	// 61 62 31 32 41 42
	fmt.Println()
}
