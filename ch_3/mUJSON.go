package main

import (
	"encoding/json"
	"fmt"
)

type SomeRecord struct {
	Name    string
	Surname string
	Tel     []SomeTelephone
}

type SomeTelephone struct {
	Mobile bool
	Number string
}

func main() {
	myRecord := SomeRecord{
		Name:    "Mihalis",
		Surname: "Tsoukalos",
		Tel: []SomeTelephone{
			{
				Mobile: true,
				Number: "1234-567",
			},
			{
				Mobile: true,
				Number: "1234-abcd",
			},
			{
				Mobile: false,
				Number: "abcd-567",
			},
		},
	}

	rec, err := json.Marshal(&myRecord)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(rec))

	var unRec SomeRecord
	err1 := json.Unmarshal(rec, &unRec)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(unRec)
}
