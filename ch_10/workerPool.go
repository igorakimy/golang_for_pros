package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type Client struct {
	id      int
	integer int
}

type Data struct {
	job    Client
	square int
}

// Создать три глобальные переменные.
var (
	size = 10
	// Буферизованные каналы clients и data используются для получения новых данных
	// клиентских запросов и записи результатов соответственно.
	clients = make(chan Client, size)
	data    = make(chan Data, size)
)

func worker(w *sync.WaitGroup) {
	// Читать канал clients.
	for c := range clients {
		square := c.integer * c.integer
		output := Data{c, square}
		// После завершения обработки результат записывается в канал data.
		data <- output
		time.Sleep(time.Second)
	}
	w.Done()
}

func makeWP(n int) {
	var w sync.WaitGroup
	// Сгенерировать необходимое количество горутин worker()
	// для обработки всех запросов.
	for i := 0; i < n; i++ {
		// Функция w.Add() вызывается из makeWP(), однако w.Done() вызывается из worker()
		// после того, как обработчик завершит свою работу.
		w.Add(1)
		go worker(&w)
	}
	w.Wait()
	close(data)
}

func create(n int) {
	// Правильно создать все запросы, используя тип Client,
	// а затем записать их в канал clients для обработки.
	for i := 0; i < n; i++ {
		c := Client{i, i}
		clients <- c
	}
	close(clients)
}

func main() {
	// Использование функции cap() для определения пропускной способности канала.
	fmt.Println("Capacity of clients:", cap(clients))
	fmt.Println("Capacity of data:", cap(data))

	if len(os.Args) != 3 {
		fmt.Println("Need #jobs and #workers!")
		os.Exit(1)
	}

	nJobs, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	nWorkers, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Вызов функции create() для имитации пользовательских запросов,
	// которые следует обработать.
	go create(nJobs)
	// Канал finished используется для блокировки программы до тех пор, пока анонимная
	// программа не закончит чтение канала data. Поэтому канал finished не нуждается
	// в конкретном типе.
	finished := make(chan interface{})
	// Для чтения канал data и вывода результатов на экран применяется анонимная горутина.
	go func() {
		for d := range data {
			fmt.Printf("Client ID: %d\tint: ", d.job.id)
			fmt.Printf("%d\tsquare: %d\n", d.job.integer, d.square)
		}
		// Хотя горутина и записывает значение true в канал finished, мы могли бы
		// записать туда значение false и получить тот же результат, который
		// разблокирует функцию main().
		finished <- true
	}()
	// Затем вызывается функция makeWP(), которая выполняет обработку запросов.
	makeWP(nWorkers)
	// Оператор <-finished не позволяет программе завершится, пока кто-то
	// не напишет что-либо в канал finished. Этот "кто-то" на самом деле
	// - анонимная горутина в main().
	fmt.Printf(": %v\n", <-finished)
}
