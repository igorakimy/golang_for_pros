package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Your are using ", runtime.Compiler, " ")
	fmt.Println("on a", runtime.GOARCH, "machine")
	fmt.Println("Using Go version", runtime.Version())
	fmt.Println("Number of CPUs:", runtime.NumCPU())
	fmt.Println("Number of Coroutines:", runtime.NumGoroutine())
}
