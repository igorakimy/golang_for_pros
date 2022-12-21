package main

import (
	"fmt"
)

func main() {
	anArray := [5]string{"RU", "FR", "BR", "CH", "UK"}
	aMap := make(map[int]string)
	for i := 0; i < len(anArray); i++ {
		aMap[i] = anArray[i]
	}
	fmt.Println(aMap)
}
