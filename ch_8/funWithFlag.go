package main

import (
	"flag"
	"fmt"
	"strings"
)

type NamesFlag struct {
	Names []string
}

func (s *NamesFlag) GetNames() []string {
	return s.Names
}

func (s *NamesFlag) String() string {
	return fmt.Sprint(s.Names)
}

// Set гарантирует, что соответствующему аргументу командной строки
// еще не присвоено значение.
func (s *NamesFlag) Set(v string) error {
	if len(s.Names) > 0 {
		return fmt.Errorf("cannot use names flag more than once")
	}

	// Разделить входные данные на отдельные аргументы.
	names := strings.Split(v, ",")
	// Сохранить аргументы в поле Names структуры NamesFlag
	for _, item := range names {
		s.Names = append(s.Names, item)
	}
	return nil
}

func main() {
	var manyNames NamesFlag
	minusK := flag.Int("k", 0, "An int")
	minusO := flag.String("o", "Mihalis", "The name")
	// Создать флаг любого типа, который соответствует интерфейсу flag.Value
	flag.Var(&manyNames, "names", "Comma-separated list")

	flag.Parse()
	fmt.Println("-k:", *minusK)
	fmt.Println("-o:", *minusO)

	for i, item := range manyNames.GetNames() {
		fmt.Println(i, item)
	}

	fmt.Println("Remaining command line arguments")
	// Срез flag.Args() содержит оставшиеся аргументы командной строки,
	// а переменная manyNames - значения аргументов командной строки из flag.Var().
	for index, val := range flag.Args() {
		fmt.Println(index, ":", val)
	}
}
