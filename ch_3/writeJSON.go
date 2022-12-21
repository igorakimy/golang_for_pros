package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type NewRecord struct {
	Name    string
	Surname string
	Tel     []NewTelephone
}

type NewTelephone struct {
	Mobile bool
	Number string
}

func saveToJSON(filename *os.File, key interface{}) {
	encodeJSON := json.NewEncoder(filename)
	err := encodeJSON.Encode(key)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	myRecord := NewRecord{
		Name:    "Mihalis",
		Surname: "Tsoukalos",
		Tel: []NewTelephone{
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
				Number: "abcc-567",
			},
		},
	}
	saveToJSON(os.Stdout, myRecord)
}
