package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

func main() {
	// Открываем утилиту /dev/random как обычный файл,
	// поскольку все в UNIX - это файлы.
	f, err := os.Open("/dev/random")
	defer f.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	var seed int64
	// Для чтения данных с системного устройства необходима функция binary.Read(),
	// которая принимает три аргумента. Значение второго из них (binary.LittleEndian)
	// указывает, что мы хотим использовать прямой порядок байтов.
	binary.Read(f, binary.LittleEndian, &seed)
	fmt.Println("Seed:", seed)
}
