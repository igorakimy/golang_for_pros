package main

import (
	"fmt"
)

func main() {
	// Определение буферизованного канала.
	numbers := make(chan int, 5)
	counter := 10

	// Попытка поместить 10 чисел в канал numbers.
	for i := 0; i < counter; i++ {
		select {
		case numbers <- i:
		default:
			fmt.Println("Not enough space for", i)
		}
	}

	// Попытка прочитать содержимое канала numbers.
	for i := 0; i < counter+5; i++ {
		select {
		// Пока в канале numbers есть что читать, будет выполняться
		// первая ветвь оператора select.
		case num := <-numbers:
			fmt.Println(num)
		// Когда канал numbers пуст, выполняется вторая ветвь.
		default:
			fmt.Println("Nothing more to be done!")
			break
		}
	}
}
