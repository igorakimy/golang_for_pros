package main

import (
	"fmt"
)

type f struct {
	XX int
	YY int
}

type g struct {
	AA string
	XX int
}

type h struct {
	F f
	G g
}

func (F f) F() {
	fmt.Println("Function F() for F")
}

func (G g) F() {
	fmt.Println("Function F() for G")
}

func main() {
	var i h
	i.F.F()
	i.G.F()

	// Function F() for F
	// Function F() for G
}
