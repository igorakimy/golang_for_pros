package main

import (
	"fmt"
)

func getPointer(n *int) {
	*n = *n * *n
}

func returnPointer(n int) *int {
	v := n * n
	return &v
}

func main() {
	// Переменные
	i := -10
	j := 25

	// Указатели
	pI := &i
	pJ := &j

	fmt.Println("pI memory:", pI)
	// pI memory: 0xc0000100b8
	fmt.Println("pJ memory:", pJ)
	// pJ memory: 0xc0000100d0
	fmt.Println("pI value:", *pI)
	// pI value: -10
	fmt.Println("pJ value:", *pJ)
	// pJ value: 25

	*pI = 123456
	*pI--
	fmt.Println("i:", i)
	// i: 123455

	getPointer(pJ)
	fmt.Println("j:", j)
	// 625
	k := returnPointer(12)
	fmt.Println(*k)
	// 144
	fmt.Println(k)
	// 0xc0000a60c0
}
