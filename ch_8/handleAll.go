package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handle(signal os.Signal) {
	fmt.Println("Received:", signal)
}

func main() {
	sigs := make(chan os.Signal, 1)
	// Обрабатывать все сигналы
	signal.Notify(sigs)

	go func() {
		for {
			sig := <-sigs
			switch sig {
			case os.Interrupt:
				handle(sig)
			case syscall.SIGTERM:
				handle(sig)
				// Очень удобно использовать один из сигналов дл выхода из программы.
				os.Exit(0)
			case syscall.SIGUSR2:
				fmt.Println("Handling syscall SIGUSR2!")
			default:
				fmt.Println("Ignoring:", sig)
			}
		}
	}()

	for {
		fmt.Printf(".")
		time.Sleep(20 * time.Second)
	}
}
