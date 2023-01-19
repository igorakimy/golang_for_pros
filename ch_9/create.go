package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	// Получить значение аргумента командной строки.
	n := flag.Int("n", 10, "Number of goroutines")
	flag.Parse()

	count := *n
	fmt.Printf("Going to create %d goroutines.\n", count)

	// На основе переданного аргумента создать определенное количество горутин.
	for i := 0; i < count; i++ {
		go func(x int) {
			fmt.Printf("%d ", x)
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println("\nExiting...")
}
