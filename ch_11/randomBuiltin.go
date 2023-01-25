package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing/quick"
	"time"
)

func main() {
	type point3D struct {
		X, Y, Z int8
		S       float32
	}
	// Создаем генератор случайных чисел.
	ran := rand.New(rand.NewSource(time.Now().Unix()))
	// Используем рефлексию, чтобы получить информацию о типе point3D.
	myValues := reflect.TypeOf(point3D{})
	// Поместить в переменную myValues некоторые случайные значения.
	x, _ := quick.Value(myValues, ran)
	fmt.Println(x)
}
