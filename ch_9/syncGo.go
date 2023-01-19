package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {
	n := flag.Int("n", 20, "Number of goroutines")
	flag.Parse()
	count := *n
	fmt.Printf("Going to create %d goroutines\n", count)

	var waitGroup sync.WaitGroup

	fmt.Printf("%#v\n", waitGroup)
	// Создаем желаемое количество горутин.
	for i := 0; i < count; i++ {
		// Каждый вызов sync.Add() увеличивает счетчик
		// в переменной sync.WaitGroup на единицу.
		// Нужно выполнять перед оператором go, чтобы предотвратить состояние гонки.
		waitGroup.Add(1)
		go func(x int) {
			// Когда горутина завершает свою работу, вызывается метод
			// sync.Done(), который уменьшает счетчик на единицу.
			defer waitGroup.Done()
			fmt.Printf("%d ", x)
		}(i)
	}
	waitGroup.Done()
	fmt.Printf("%#v\n", waitGroup)
	// Вызов sync.Wait() блокируется до тех пор, пока значение
	// счетчика в соответствующей переменной sync.WaitGroup() не станет
	// равным нулю, что дает возможность всем горутинам завершить работу.
	waitGroup.Wait()
	fmt.Println("\nExiting...")
}
