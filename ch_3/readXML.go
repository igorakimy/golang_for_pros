package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Record2 struct {
	Name    string
	Surname string
	Tel     []Telephone2
}

type Telephone2 struct {
	Mobile bool
	Number string
}

func loadFromXML(filename string, key interface{}) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}

	decodeXML := xml.NewDecoder(in)
	err = decodeXML.Decode(key)
	if err != nil {
		return err
	}
	_ = in.Close()
	return nil
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a filename!")
		return
	}

	filename := arguments[1]

	var myRecord Record2
	err := loadFromXML(filename, &myRecord)
	if err == nil {
		fmt.Println("XML:", myRecord)
	} else {
		fmt.Println(err)
	}
}
