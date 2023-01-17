package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func lineByLine(file string) error {
	// Создать пустую переменную для ошибки
	var err error
	// Открыть файл
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	// Закрыть файл, после выполнения функции.
	defer f.Close()

	// Создать нового читателя.
	r := bufio.NewReader(f)
	for {
		// Использовать читателя для строчного чтения входного файла.
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			break
		}
		fmt.Print(line)
	}
	return nil
}

func main() {
	// Обработка аргументов командной строки.
	flag.Parse()
	// Если аргументы не были переданы.
	if len(flag.Args()) == 0 {
		// Выводим ошибку.
		fmt.Printf("usage: byLine <file1> [<file2> ...]\n")
		return
	}

	// Обходим аргументы командной строки, которые являются именами файлов.
	for _, file := range flag.Args() {
		// Считать файл строка за строкой.
		err := lineByLine(file)
		if err != nil {
			fmt.Println(err)
		}
	}
}
