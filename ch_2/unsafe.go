package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var value int64 = 5
	var p1 = &value
	var p2 = (*int32)(unsafe.Pointer(p1))

	// Разыменование указателя
	fmt.Println("*p: ", *p1)
	fmt.Println("*p2: ", *p2)
	// Изменение значения через указатель
	*p1 = 5434123412312431212
	fmt.Println(value)
	fmt.Println("*p2: ", *p2)
	*p1 = 54341234
	fmt.Println(value)
	fmt.Println("*p2: ", *p2)
}
