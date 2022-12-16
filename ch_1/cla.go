package main

import (
	"fmt"
	"os"
	"strconv"
)

// Найти минимальный и максимальный из своих аргументов командной строки
func main() {
	// Проверить, есть ли аргументы командной строки
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more floats.")
		os.Exit(1)
	}

	arguments := os.Args

	// Спарсить первый аргумент командной строки
	min, _ := strconv.ParseFloat(arguments[1], 64)
	max, _ := strconv.ParseFloat(arguments[1], 64)

	// Спарсить остальные аргументы командной строки для нахождения
	// минимального и максимального значения среди всех аргументов
	for i := 2; i < len(arguments); i++ {
		n, _ := strconv.ParseFloat(arguments[i], 64)
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
