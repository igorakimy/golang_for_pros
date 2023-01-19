package main

import (
	"fmt"
	"time"
)

func writeToChannel2(c chan int, x int) {
	fmt.Println("1", x)
	c <- x
	close(c)
	fmt.Println("2", x)
}

func main() {
	c := make(chan int)
	go writeToChannel2(c, 10)
	// Предоставить время для записи в канал.
	time.Sleep(1 * time.Second)
	fmt.Println("Read:", <-c)
	// Предоставить время для чтения из канала.
	time.Sleep(1 * time.Second)

	// Прием, позволяющий определить, открыт ли канал.
	_, ok := <-c
	if ok {
		fmt.Println("Channel is open!")
	} else {
		fmt.Println("Channel is closed!")
	}
}
