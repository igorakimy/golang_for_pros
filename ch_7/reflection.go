package main

import (
	"fmt"
	"os"
	"reflect"
)

type a struct {
	X int
	Y float64
	Z string
}

type b struct {
	F int
	G int
	H string
	I float64
}

func main() {
	x := 100
	xRefl := reflect.ValueOf(&x).Elem()
	xType := xRefl.Type()
	fmt.Printf("The type of x is %s.\n", xType)

	A := a{100, 200.12, "Struct a"}
	B := b{1, 2, "Struct b", -1.2}
	var r reflect.Value
	arguments := os.Args
	if len(arguments) == 1 {
		r = reflect.ValueOf(&A).Elem()
	} else {
		r = reflect.ValueOf(&B).Elem()
	}

	iType := r.Type()
	fmt.Printf("i Type: %s\n", iType)
	fmt.Printf("The %d fields of %s are: \n", r.NumField(), iType)
	for i := 0; i < r.NumField(); i++ {
		fmt.Printf("Field name: %s ", iType.Field(i).Name)
		fmt.Printf("with type: %s ", r.Field(i).Type())
		fmt.Printf("and value %v\n", r.Field(i).Interface())
	}

	// $ go run reflection.go 1
	// The type of x is int.
	// i Type: main.b
	// The 4 fields of main.b are:
	// Field name: F with type: int and value 1
	// Field name: G with type: int and value 2
	// Field name: H with type: string and value Struct b
	// Field name: I with type: float64 and value -1.2

	// $ go run reflection.go
	// The type of x is int.
	// i Type: main.a
	// The 3 fields of main.a are:
	// Field name: X with type: int and value 100
	// Field name: Y with type: float64 and value 200.12
	// Field name: Z with type: string and value Struct a
}
