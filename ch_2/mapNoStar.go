package main

import (
	"runtime"
)

func main() {
	var N = 40000000
	myMap := make(map[int]int)
	// Помещение в хэш-таблицу целочисленных значений
	for i := 0; i < N; i++ {
		value := int(i)
		myMap[value] = value
	}
	runtime.GC()
	_ = myMap[0]
}
