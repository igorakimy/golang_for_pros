package main

import (
	"fmt"
)

func main() {
	c1 := 12 + 1i
	c2 := complex(5, 7)
	fmt.Printf("Type of c1: %T\n", c1)
	// Type of c1: complex128
	fmt.Printf("Type of c2: %T\n", c2)
	// Type of c2: complex128

	var c3 complex64 = complex64(c1 + c2)
	fmt.Println("c3:", c3)
	// c3: (17+8i)
	fmt.Printf("Type of c3: %T\n", c3)
	// Type of c3: complex64

	cZero := c3 - c3
	fmt.Println("cZero:", cZero)
	// cZero: (0+0i)

	x := 12
	k := 5
	fmt.Println(x)
	// 12
	fmt.Printf("Type of x: %T\n", x)
	// Type of x: int

	div := x / k
	fmt.Println("div", div)
	// div 2

	var m, n float64
	m = 1.223
	fmt.Println("m, n:", m, n)
	// m, n: 1.223 0

	y := 4 / 2.3
	fmt.Println("y:", y)
	// y: 1.7391304347826086

	divFloat := float64(x) / float64(y)
	fmt.Println("divFloat", divFloat)
	// divFloat 6.9
	fmt.Printf("Type of divFloat: %T\n", divFloat)
	// Type of divFloat: float64

}
