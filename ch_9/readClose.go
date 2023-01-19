package main

import (
	"fmt"
)

func main() {
	willClose := make(chan int, 10)

	willClose <- -1
	willClose <- 0
	willClose <- 2

	<-willClose
	<-willClose
	<-willClose

	// Закрыть канал willClose.
	close(willClose)
	// Попытаться прочитать данные из закрытого канала.
	read := <-willClose
	// При чтении данных из закрытого канала возвращается нулевое значение
	// соответствующего типа.
	fmt.Println(read)
}
