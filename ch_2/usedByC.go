package main

import "C"

import (
	"fmt"
)

func main() {}

//export PrintMessage
func PrintMessage() {
	fmt.Println("A Go function!")
}

//export Multiply
func Multiply(a, b int) int {
	return a * b
}
