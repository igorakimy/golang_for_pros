package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func gen(min, max int, createNumber chan int, end chan bool) {
	for {
		// Оператор select не выполняется последовательно, т.к. все его каналы
		// проверяются одновременно.
		// Если ни один из каналов, указанных в операторе select не доступен, то
		// оператор select будет заблокирован, пока не освободится один из каналов.
		// Если доступны сразу несколько каналов, то среда выполнения Go сделает
		// случайный выбор из набора этих доступных каналов.
		select {
		case createNumber <- rand.Intn(max-min) + min:
		case <-end:
			close(end)
			return
		// time.After() ожидает истечения заданного интервала, после чего передает
		// значение текущего времени во возвращаемому каналу - это разблокирует
		// оператор select, если все остальные каналы по какой-либо причине
		// окажутся заблокированными.
		case <-time.After(4 * time.Second):
			fmt.Println("\ntime.After()!")
		}
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	createNumber := make(chan int)
	end := make(chan bool)
	if len(os.Args) != 2 {
		fmt.Println("Please give me an integer!")
		return
	}

	n, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("Going to create %d random numbers.\n", n)
	go gen(0, 2*n, createNumber, end)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", <-createNumber)
	}

	// Предоставить функциям time.After() и gen() достаточно времени,
	// чтобы обработать, вернуть результат и таким образом активировать
	// соответствующую ветвь оператора select.
	time.Sleep(5 * time.Second)
	fmt.Println("Exiting...")
	// Активировать ветвь case <-end оператора select в gen().
	end <- true
}
