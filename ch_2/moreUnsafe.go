package main

import (
	"fmt"
	"unsafe"
)

func main() {
	array := [...]int{0, 1, -2, 3, 4}
	// Переменная pointer указывает на адрес в памяти первого элемента массива
	pointer := &array[0]
	fmt.Print(*pointer, " ")
	// Переменная pointer преобразовывается в unsafe.Pointer(), потом в uintptr
	memoryAddress := uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])

	for i := 0; i < len(array)-1; i++ {
		pointer = (*int)(unsafe.Pointer(memoryAddress))
		// *pointer разыменовывает указатель и возвращает сохраненное значение
		fmt.Print(*pointer, " ")
		memoryAddress = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
	}

	fmt.Println()
	pointer = (*int)(unsafe.Pointer(memoryAddress))
	fmt.Print("One more: ", *pointer, " ")
	memoryAddress = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
	fmt.Println()
}
