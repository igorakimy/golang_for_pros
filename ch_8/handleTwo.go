package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handleSignal(signal os.Signal) {
	fmt.Println("handleSignal() Caught:", signal)
}

func main() {
	// Определить канал с именем sigs, по которому будут передаваться данные.
	sigs := make(chan os.Signal, 1)
	// Указать на интересующие нас сигналы
	signal.Notify(sigs, os.Interrupt, syscall.SIGINT)
	// Реализация анонимной функции, которая выполняется как горутина
	// и включается тогда, когда поступает любой из интересующих нас сигналов.
	go func() {
		for {
			sig := <-sigs
			switch sig {
			case os.Interrupt:
				fmt.Println("Caught:", sig)
			case syscall.SIGINT:
				handleSignal(sig)
				return
			}
		}
	}()

	for {
		fmt.Printf(".")
		// Вызов time.Sleep() не дает программе завершиться, поскольку она не
		// выполняет никакой полезной функции.
		time.Sleep(20 * time.Second)
	}
}
