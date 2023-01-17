package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"os"
	"strings"
)

type MyElement struct {
	Name    string
	Surname string
	Id      string
}

var DATAFILE = "dataFile.gob"
var DATA = make(map[string]MyElement)

func save() error {
	fmt.Println("Saving", DATAFILE)
	// Удалить существующий файл данных.
	err := os.Remove(DATAFILE)
	if err != nil {
		fmt.Println(err)
	}

	// Создать новый файл данных.
	saveTo, err := os.Create(DATAFILE)
	if err != nil {
		fmt.Println("Cannot create", DATAFILE)
		return err
	}
	defer func(saveTo *os.File) {
		err := saveTo.Close()
		if err != nil {
			fmt.Println("Error on close file")
		}
	}(saveTo)

	encoder := gob.NewEncoder(saveTo)
	err = encoder.Encode(DATA)
	if err != nil {
		fmt.Println("Cannot save to", DATAFILE)
		return err
	}
	return nil
}

func load() error {
	fmt.Println("Loading", DATAFILE)
	loadFrom, err := os.Open(DATAFILE)
	defer loadFrom.Close()
	if err != nil {
		fmt.Println("Empty key/value store!")
		return err
	}

	decoder := gob.NewDecoder(loadFrom)
	decoder.Decode(&DATA)
	return nil
}

func ADD(k string, n MyElement) bool {
	if k == "" {
		return false
	}

	if LOOKUP(k) == nil {
		DATA[k] = n
		return true
	}
	return false
}

func DELETE(k string) bool {
	if LOOKUP(k) != nil {
		delete(DATA, k)
		return true
	}
	return false
}

func LOOKUP(k string) *MyElement {
	_, ok := DATA[k]
	if ok {
		n := DATA[k]
		return &n
	}
	return nil
}

func CHANGE(k string, n MyElement) bool {
	DATA[k] = n
	return true
}

func PRINT() {
	for k, d := range DATA {
		fmt.Printf("key: %s value: %v\n", k, d)
	}
}

func main() {
	err := load()
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.TrimSpace(text)
		tokens := strings.Fields(text)

		switch len(tokens) {
		case 0:
			continue
		case 1:
			for i := 0; i < 4; i++ {
				tokens = append(tokens, "")
			}
		case 2:
			for i := 0; i < 3; i++ {
				tokens = append(tokens, "")
			}
		case 3:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 4:
			tokens = append(tokens, "")
		}

		switch tokens[0] {
		case "PRINT":
			PRINT()
		case "STOP":
			err = save()
			if err != nil {
				fmt.Println(err)
			}
			return
		case "DELETE":
			if !DELETE(tokens[1]) {
				fmt.Println("Delete operation failed!")
			}
		case "ADD":
			n := MyElement{tokens[2], tokens[3], tokens[4]}
			if !ADD(tokens[1], n) {
				fmt.Println("Add operation failed!")
			}
		case "LOOKUP":
			n := LOOKUP(tokens[1])
			if n != nil {
				fmt.Printf("%v\n", *n)
			}
		case "CHANGE":
			n := MyElement{tokens[2], tokens[3], tokens[4]}
			if !CHANGE(tokens[1], n) {
				fmt.Println("Update operation failed!")
			}
		default:
			fmt.Println("Unknown command - please try again!")
		}
	}
}
