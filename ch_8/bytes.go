package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	// Создаем переменную типа bytes.Buffer
	var buffer bytes.Buffer
	// Записываем в переменную данные
	buffer.Write([]byte("This is"))
	fmt.Fprintf(&buffer, " a string!\n")
	// При первом вызове WriteTo() будет выведено содержимое переменной buffer.
	buffer.WriteTo(os.Stdout)
	// Однако при втором вызове ничего не выведется,
	// поскольку переменная buffer будет очищена.
	buffer.WriteTo(os.Stdout)

	// Метод Reset() обнуляет переменную buffer.
	buffer.Reset()
	// А метод Write() снова запишет в нее данные.
	buffer.Write([]byte("Mastering Go!"))
	// Создать новый объект Reader интерфейса io.Reader.
	r := bytes.NewReader([]byte(buffer.String()))
	fmt.Println(buffer.String())
	for {
		b := make([]byte, 3)
		// Считать из переменной buffer данные в буфер.
		n, err := r.Read(b)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Read %s Bytes: %d", b, n)
	}
}
