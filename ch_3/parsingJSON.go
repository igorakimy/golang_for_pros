package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a filename!")
		return
	}

	filename := arguments[1]

	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	var parsedData map[string]interface{}
	err = json.Unmarshal([]byte(fileData), &parsedData)
	if err != nil {
		return
	}

	for key, value := range parsedData {
		fmt.Println("key:", key, "value:", value)
	}
}
