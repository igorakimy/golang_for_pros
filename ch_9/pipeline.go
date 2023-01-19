package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// CLOSEA - Переменная для отслеживания состояния первого канала.
var CLOSEA = false

var DATA = make(map[int]bool)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func first(min, max int, out chan<- int) {
	for {
		// Продолжать выполнение цикла до тех пор, пока логическая
		// переменная CLOSEA не станет равной true.
		if CLOSEA {
			// После чего функция закроет свой канал out.
			close(out)
			return
		}
		out <- random(min, max)
	}
}

func second(out chan<- int, in <-chan int) {
	// Считать данные из канала in и автоматически закрыть его,
	// когда будут прочитаны все данные.
	for x := range in {
		fmt.Print(x, " ")
		_, ok := DATA[x]
		if ok {
			// Если случайное число уже существует в карте,
			// то глобальной переменной CLOSEA присваивается значение true.
			// Соответственно, прекращается дальнейшая передача других
			// чисел в канал out. После этого функция закрывает канал out.
			CLOSEA = true
		} else {
			// Если число еще не существует в карте, то добавить его.
			DATA[x] = true
			// Отправить число в канал out.
			out <- x
		}
	}
	fmt.Println()
	close(out)
}

func third(in <-chan int) {
	var sum int
	// Считать данные из канала in. Когда функция second() закрывает этот канал,
	// цикл for прекращает получать данные и функция выводит результат.
	for x2 := range in {
		sum = sum + x2
	}
	fmt.Printf("The sum of the random numbers is %d\n", sum)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Need two integer parameters!")
	}

	n1, _ := strconv.Atoi(os.Args[1])
	n2, _ := strconv.Atoi(os.Args[2])

	if n1 > n2 {
		fmt.Printf("%d should be smaller than %d", n1, n2)
		return
	}

	rand.Seed(time.Now().UnixNano())
	A := make(chan int)
	B := make(chan int)

	go first(n1, n2, A)
	go second(B, A)
	third(B)
}
