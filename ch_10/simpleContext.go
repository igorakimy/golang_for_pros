package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"
)

func func1(t int) {
	// Инициализация пустого объекта Context.
	c1 := context.Background()
	// Функция context.WithCancel() использует существующий объект Context
	// и создает его потомка с возможностью аннулирования.
	// Также данная функция создает канал Done, который можно закрыть либо при
	// вызове функции cancel(), либо когда закрывается канал Done родительского контекста.
	c1, cancel := context.WithCancel(c1)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	// При вызове функции Done() происходить аннулирование контекста.
	case <-c1.Done():
		fmt.Println("func1():", c1.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("func1():", r)
	}
	return
}

func func2(t int) {
	c2 := context.Background()
	// По окончании времени ожидания автоматически вызывается функция cancel().
	c2, cancel := context.WithTimeout(c2, time.Duration(t)*time.Second)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c2.Done():
		fmt.Println("func2():", c2.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("func2():", r)
	}
	return
}

func func3(t int) {
	c3 := context.Background()
	deadline := time.Now().Add(time.Duration(2*t) * time.Second)
	// Задание крайнего срока окончания операции. По истечению этого времени
	// автоматически вызывается функция cancel().
	c3, cancel := context.WithDeadline(c3, deadline)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c3.Done():
		fmt.Println("func3():", c3.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("func3():", r)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Need a delay!")
		return
	}

	delay, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Delay:", delay)

	func1(delay)
	func2(delay)
	func3(delay)
}
