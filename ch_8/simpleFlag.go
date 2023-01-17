package main

import (
	"flag"
	"fmt"
)

func main() {
	// Определить логический элемент командной строки с именем k, значение по
	// умолчанию которого равно true.
	minusK := flag.Bool("k", true, "k flag")
	// Добавить в программу поддержку целочисленного аргумента командной строки.
	minusO := flag.Int("O", 1, "0")
	// После определения всех аргументов командной строки всегда
	// необходимо вызывать метод flag.Parse().
	flag.Parse()

	// Получить значения аргументов командной строки можно таким образом.
	valueK := *minusK
	valueO := *minusO
	valueO++

	fmt.Println("-k:", valueK)
	fmt.Println("-O:", valueO)

	// $ go run simpleFlag.go -O 100
	// -k: true
	// -O: 101

	// $ go run simpleFlag.go -O=100
	// -k: true
	// -O: 101

	// $ go run simpleFlag.go -O=100 -k
	// -k: true
	// -O: 101

	// $ go run simpleFlag.go -O=100 -k false
	// -k: true
	// -O: 101

	// $ go run simpleFlag.go -O=100 -k=false
	// -k: false
	// -O: 101
}
