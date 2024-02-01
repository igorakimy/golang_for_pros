package main

import "fmt"

type Secret struct {
	SecretValue string
}

type Entry struct {
	F1 int
	F2 string
	F3 Secret
}

func TestStruct(x interface{}) {
	// переключатель типа
	switch T := x.(type) {
	case Secret:
		fmt.Println("Secret type")
	case Entry:
		fmt.Println("Entry type")
	default:
		fmt.Printf("Not suported type: %T\n", T)
	}
}

func Learn(x interface{}) {
	switch T := x.(type) {
	default:
		fmt.Printf("Data type: %T\n", T)
	}
}

func main() {
	A := Entry{100, "F2", Secret{"myPassword"}}
	TestStruct(A)
	TestStruct(A.F3)
	TestStruct("A string")

	Learn(12.23)
	Learn('€')
	Learn('\t')
}
