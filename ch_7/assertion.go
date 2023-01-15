package main

import (
	"fmt"
)

func main() {
	// Объявляем переменную myInt с динамическим типом int и значением 123
	var myInt interface{} = 123

	// Использование утверждения типа для проверки интерфейса переменной myInt
	// для типа int
	k, ok := myInt.(int)
	if ok {
		fmt.Println("Success:", k)
	}

	// Использование утверждения типа для проверки интерфейса переменной myInt
	// для типа float64
	v, ok := myInt.(float64)
	if ok {
		fmt.Println(v)
	} else {
		// Поскольку переменная myInt не содержит значения типа float64, операция
		// утверждения типа myInt.(float64) завершится сбоем. Переменная ok не даст
		// программе уйти в панику.
		fmt.Println("Failed without panicking!")
	}

	i := myInt.(int)
	fmt.Println("No checking:", i)
	// No checking: 123

	// Здесь утверждение типа вызывает панику, т.к. основное значение myInt не
	// является логическим (bool)
	j := myInt.(bool)
	fmt.Println(j)
}
