package main

import (
	"fmt"
	"time"
)

func A(a, b chan struct{}) {
	// Функция A() заблокирована каналом, который храниться в переменной a.
	// Как только этот канал будет разблокирован в main(), функция A() начнет работу.
	<-a
	fmt.Println("A()!")
	time.Sleep(time.Second)
	// В конце, функция A() закроет канал b, тем самым разблокировав другую функцию - B().
	close(b)
}

func B(a, b chan struct{}) {
	// Функция B() блокируется, пока не будет закрыт канал a.
	<-a
	fmt.Println("B()!")
	close(b)
}

func C(a chan struct{}) {
	// Функция C() заблокирована и ожидает закрытия канала a, чтобы начать работу.
	<-a
	fmt.Println("C()!")
}

func main() {
	x := make(chan struct{})
	y := make(chan struct{})
	z := make(chan struct{})

	go C(z)
	go A(x, y)
	go C(z)
	go B(y, z)
	go C(z)

	close(x)
	time.Sleep(3 * time.Second)
}
