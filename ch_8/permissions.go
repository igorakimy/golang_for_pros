package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Printf("usage: permissions filename\n")
		return
	}

	filename := os.Args[1]
	// Вызов os.Stat() возвращает структуру с большим количеством данных.
	info, _ := os.Stat(filename)
	// Поскольку нас интересуют только права доступа к файлу, мы вызываем метод Mode().
	mode := info.Mode()
	fmt.Println(filename, "mode is", mode.String()[1:10])
}
