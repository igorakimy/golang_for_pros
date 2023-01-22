package main

import (
	"fmt"
	"runtime"
)

func getGOMAXPROCS() int {
	// Получить значение кол-ва логических процессоров на машине.
	return runtime.GOMAXPROCS(0)
}

func main() {
	fmt.Printf("GOMAXPROCS: %d\n", getGOMAXPROCS())
}
