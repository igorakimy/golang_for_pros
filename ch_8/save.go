package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	s := []byte("Data to write\n")

	// Запись в файл при помощи fmt.Fprintf().
	f1, err := os.Create("f1.txt")
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	defer f1.Close()
	fmt.Fprintf(f1, string(s))

	// Запись в файл при помощи WriteString().
	f2, err := os.Create("f2.txt")
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	defer f2.Close()
	n, err := f2.WriteString(string(s))
	fmt.Printf("wrote %d bytes\n", n)

	// Запись в файл при помощи bufio.WriteString().
	f3, err := os.Create("f3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	w := bufio.NewWriter(f3)
	n, err = w.WriteString(string(s))
	fmt.Printf("wrote %d bytes\n", n)
	w.Flush()

	// Запись в файл при помощи os.WriteFile().
	f4 := "f4.txt"
	err = os.WriteFile(f4, s, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Запись в файл при помощи io.WriteString().
	f5, err := os.Create("f5.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	n, err = io.WriteString(f5, string(s))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("wrote %d bytes\n", n)
}
