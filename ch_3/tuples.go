package main

import (
	"fmt"
)

func retThree(x int) (int, int, int) {
	return 2 * x, x * x, -x
}

func main() {
	fmt.Println(retThree(10))
	// 20 100 -10
	n1, n2, n3 := retThree(20)
	fmt.Println(n1, n2, n3)
	// 40 400 -20

	n1, n2 = n2, n1
	fmt.Println(n1, n2, n3)
	// 400 40 -20

	x1, x2, x3 := n1*2, n1*n1, -n1
	fmt.Println(x1, x2, x3)
	// 800 160000 -400
}
