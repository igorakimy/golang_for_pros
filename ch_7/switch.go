package main

import (
	"fmt"
)

type square2 struct {
	X float64
}

type circle2 struct {
	R float64
}

type rectangle struct {
	X float64
	Y float64
}

func tellInterface(x interface{}) {
	switch v := x.(type) {
	case square2:
		fmt.Println("This is a square!")
	case circle2:
		fmt.Printf("%v is a circle!\n", v)
	case rectangle:
		fmt.Println("This is a rectangle!")
	default:
		fmt.Printf("Unknown type %T!\n", v)
	}
}

func main() {
	x := circle2{R: 10}
	tellInterface(x)
	y := rectangle{X: 4, Y: 1}
	tellInterface(y)
	z := square2{X: 4}
	tellInterface(z)
	tellInterface(10)

	// {10} is a circle!
	// This is a rectangle!
	// This is a square!
	// Unknown type int!
}
