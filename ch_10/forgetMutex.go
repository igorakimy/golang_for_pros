package main

import (
	"fmt"
	"sync"
)

var mx sync.Mutex

func function() {
	mx.Lock()
	fmt.Println("Locked!")
}

func main() {
	var w sync.WaitGroup

	go func() {
		defer w.Done()
		function()
	}()
	w.Add(1)

	go func() {
		defer w.Done()
		function()
	}()
	w.Add(1)

	w.Wait()
}
