package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func printFile(filename string) error {
	// Открываем файл.
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	// Создаем сканер файла.
	scanner := bufio.NewScanner(f)
	// Сканируем файл.
	for scanner.Scan() {
		// Записываем считанную строку в поток вывода
		io.WriteString(os.Stdout, scanner.Text())
		io.WriteString(os.Stdout, "\n")
	}
	return nil
}

func main() {
	filename := ""
	// Получаем аргументы командной строки.
	arguments := os.Args
	// Если аргументы не были переданы
	if len(arguments) == 1 {
		// Копируем содержимое потока ввода в поток вывода.
		_, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	// Обходим кол-во переданных аргументов.
	for i := 1; i < len(arguments); i++ {
		// Получаем имя файла(которое и является самим аргументом)
		filename = arguments[i]
		// Выводим содержимое файла в поток вывода.
		err := printFile(filename)
		if err != nil {
			fmt.Println(err)
		}
	}
}
