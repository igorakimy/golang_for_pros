package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var DATA2 = make(map[int]bool)
var signal chan struct{}

func random2(min, max int) int {
	return rand.Intn(max-min) + min
}

func first2(min, max int, out chan<- int) {
	for {
		select {
		case <-signal:
			close(out)
			return
		case out <- random2(min, max):
		}
	}
}

func second2(out chan<- int, in <-chan int) {
	for x := range in {
		_, ok := DATA2[x]
		if ok {
			signal <- struct{}{}
			fmt.Print(x, " ")
		} else {
			DATA2[x] = true
			out <- x
		}
	}
	fmt.Println()
	close(out)
}

func third2(in <-chan int) {
	var sum int
	for x2 := range in {
		sum = sum + x2
	}
	fmt.Printf("The sum of the random numbers is %d\n", sum)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Need two integer parameters!")
	}

	n1, _ := strconv.Atoi(os.Args[1])
	n2, _ := strconv.Atoi(os.Args[2])

	if n1 > n2 {
		fmt.Printf("%d should be smaller than %d", n1, n2)
		return
	}

	signal = make(chan struct{})
	rand.Seed(time.Now().UnixNano())
	A := make(chan int)
	B := make(chan int)

	go first2(n1, n2, A)
	go second2(B, A)
	third2(B)
}
