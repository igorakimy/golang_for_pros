package main

import (
	"fmt"
	"math/rand"
	"time"
)

func add(c chan int) {
	sum := 0
	t := time.NewTimer(time.Second)

	for {
		select {
		case input := <-c:
			sum += input
		// Здесь блокируется канал C таймера t на время, указанное в вызове time.NewTimer().
		// По истечении времени таймер отправляет значение в канал t.C,
		// чем инициирует выполнение соответствующей ветви оператора select
		// - он присваивает каналу c значение nil и выводит сумму на экран.
		case <-t.C:
			c = nil
			fmt.Println(sum)
		}
	}
}

func send(c chan int) {
	// Генерировать случайные числа и отправлять их в канал до тех пор, пока канал открыт.
	for {
		c <- rand.Intn(10)
	}
}

func main() {
	c := make(chan int)
	go add(c)
	go send(c)

	time.Sleep(3 * time.Second)
}
