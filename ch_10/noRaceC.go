package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

var aMutex sync.Mutex

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Give me a natural number!")
		return
	}

	numGR, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup
	var i int
	k := make(map[int]int)
	k[1] = 12

	for i = 0; i < numGR; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			// Если не добавлять мьютекс блокировки, то будет ошибка конкурентной
			// записи в карту.
			aMutex.Lock()
			k[j] = j
			aMutex.Unlock()
		}(i)
	}

	wg.Wait()
	k[2] = 10
	fmt.Printf("k = %#v\n", k)
}
