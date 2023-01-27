package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s URL\n", filepath.Base(os.Args[0]))
	}

	URL := os.Args[1]
	// Метод http.Get() возвращает структуру http.Response.
	data, err := http.Get(URL)

	if err != nil {
		fmt.Println(err)
		return
	} else {
		defer data.Body.Close()
		// Получить содержимое поля Body структуры http.Response и записать
		// в стандартный поток вывода.
		_, err := io.Copy(os.Stdout, data.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
