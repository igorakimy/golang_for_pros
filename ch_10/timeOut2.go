package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func timeout(w *sync.WaitGroup, t time.Duration) bool {
	temp := make(chan int)
	go func() {
		defer close(temp)
		time.Sleep(5 * time.Second)
		// w.Wait() заставить функцию timeout() бесконечно ожидать соответствующую
		// функцию sync.Done(), чтобы завершить работу.
		w.Wait()
	}()

	select {
	// Когда w.Wait() закончит работу, будет выполнена первая ветвь оператора select.
	case <-temp:
		return false
	case <-time.After(t):
		return true
	}
}

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Need a time duration!")
		return
	}

	var w sync.WaitGroup
	w.Add(1)
	t, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Функция time.Duration() преобразует целочисленное значение
	// в переменную time.Duration, которая будет использована позже.
	duration := time.Duration(int32(t)) * time.Millisecond
	fmt.Printf("Timeout period is %s\n", duration)

	if timeout(&w, duration) {
		fmt.Println("Timed out!")
	} else {
		fmt.Println("OK!")
	}

	w.Done()
	if timeout(&w, duration) {
		fmt.Println("Timed out!")
	} else {
		fmt.Println("OK!")
	}
}
