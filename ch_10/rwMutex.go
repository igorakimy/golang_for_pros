package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var Password = secret{password: "myPassword"}

type secret struct {
	RWM      sync.RWMutex
	M        sync.Mutex
	password string
}

func Change(c *secret, pass string) {
	c.RWM.Lock()
	fmt.Println("LChange")
	time.Sleep(10 * time.Second)
	// Функция Change() изменяет общую переменную.
	c.password = pass
	c.RWM.Unlock()
}

func show(c *secret) string {
	// В функции show() используется RLock() и RUnlock(), поскольку ее критический
	// раздел нужен для чтения общей переменной.
	c.RWM.RLock()
	fmt.Println("show")
	time.Sleep(3 * time.Second)
	defer c.RWM.RUnlock()
	return c.password
}

func showWithLock(c *secret) string {
	// Отличие функции showWithLock() от show() лишь в том, что в ней применятся
	// монопольная блокировка для чтения данных. Следовательно, только одна функция
	// showWithLock() может прочитать поле password структуры secret.
	c.RWM.Lock()
	fmt.Println("showWithLock")
	time.Sleep(3 * time.Second)
	defer c.RWM.Unlock()
	return c.password
}

func main() {
	start := time.Now()
	var showFunction func(c *secret) string
	if len(os.Args) != 2 {
		fmt.Println("Using sync.RWMutex!")
		showFunction = show
	} else {
		fmt.Println("Using sync.Mutex")
		showFunction = showWithLock
	}

	var waitGroup sync.WaitGroup

	fmt.Println("Pass:", showFunction(&Password))

	for i := 0; i < 15; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			fmt.Println("Go Pass:", showFunction(&Password))
		}()
	}

	go func() {
		waitGroup.Add(1)
		defer waitGroup.Done()
		Change(&Password, "123456")
	}()

	waitGroup.Wait()
	fmt.Println("Pass:", showFunction(&Password))
	fmt.Println(time.Now().Sub(start))
}
