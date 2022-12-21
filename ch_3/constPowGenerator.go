package main

import (
	"fmt"
)

type Power4 int

const (
	po_1 Power4 = 4 << iota
	_
	po_2
	_
	po_3
	_
	po_4
)

func main() {
	fmt.Println("4^1", po_1)
	fmt.Println("4^2", po_2)
	fmt.Println("4^3", po_3)
	fmt.Println("4^4", po_4)
}
