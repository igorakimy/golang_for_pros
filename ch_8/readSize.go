package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func readSize(f *os.File, size int) []byte {
	// Создать буфер в виде байтового среза указанной длины.
	buffer := make([]byte, size)

	// Считать данные из файла в буфер.
	n, err := f.Read(buffer)
	// Когда файл подойдет к концу, вернуть nil.
	if err == io.EOF {
		return nil
	}

	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Вернуть срез буфера размером с кол-во прочитанных байт.
	return buffer[:n]
}

func main() {
	// Получить аргументы командной строки.
	arguments := os.Args
	// Если кол-во аргументов меньше, чем 3.
	if len(arguments) != 3 {
		// Вывести на экран сообщение.
		fmt.Println("<buffer size> <filename>")
		return
	}

	// Получить второй аргумент командной строки, который является размером буфера
	// и привести его к числовому типу, т.к. все аргументы изначально имеют строковый тип.
	bufferSize, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Получить третий аргумент командной строки, который является именем файла.
	file := os.Args[2]
	// Открыть файл.
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Отложить закрытие файла.
	defer f.Close()

	// В бесконечном цикле.
	for {
		// Считывать из указанного файла определенное кол-во байт за раз.
		readData := readSize(f, bufferSize)
		if readData != nil {
			// Если данные были считаны, то вывести их,
			// приведя байтовый срез в строковый формат.
			fmt.Println(string(readData))
		} else {
			// Иначе, прекратить считывание и выйти из цикла.
			break
		}
	}
}
