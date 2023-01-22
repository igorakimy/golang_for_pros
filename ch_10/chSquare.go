package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var times int

func f1(cc chan chan int, f chan bool) {
	// Создаем обычный канал типа int.
	c := make(chan int)
	// Передаем его каналу каналов.
	cc <- c
	defer close(c)

	sum := 0
	// Получить возможность читать данные из обычного канала типа int
	// или выходить из функции, используя сигнальный канал f.
	select {
	case x := <-c:
		// Прочитав одно значение из канала c, мы запускаем цикл, вычисляющий сумму
		// всех целых чисел от 0 до того числа, которое было прочитано.
		for i := 0; i <= x; i++ {
			sum += i
		}
		// Затем отправляем вычисленное значение в канал c.
		c <- sum
	case <-f:
		return
	}
}

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Need just one integer argument!")
		return
	}

	times, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Создать переменную канала каналов.
	cc := make(chan chan int)

	for i := 1; i < times+1; i++ {
		// Канал f является сигнальным каналом для завершения горутины,
		// когда вся работа закончена.
		f := make(chan bool)
		go f1(cc, f)
		// Инструкция ch := <-cc позволяет получить из переменной канала каналов
		// обычный канал, чтобы передать туда значение типа int.
		ch := <-cc
		ch <- i
		// Считать данные из канала.
		for sum := range ch {
			fmt.Print("Sum(", i, ")=", sum)
		}
		fmt.Println()
		time.Sleep(time.Second)
		close(f)
	}
}
