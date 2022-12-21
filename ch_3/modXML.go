package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

func main() {
	// Описать структуру адреса
	type Address struct {
		City, Country string
	}

	// Описать структуру сотрудника
	type Employee struct {
		XMLName   xml.Name `xml:"employee"`
		Id        int      `xml:"id,attr"`
		FirstName string   `xml:"name>first"`
		LastName  string   `xml:"name>last"`
		Initials  string   `xml:"name>initials"`
		Height    float32  `xml:"height,omitempty"`
		Address
		Comment string `xml:",comment"`
	}

	// Создать структуру сотрудника
	r := &Employee{
		Id:        7,
		FirstName: "Mihalis",
		LastName:  "Tsoukalos",
		Initials:  "MIT",
	}

	r.Comment = "Technical Writer + DevOps"
	r.Address = Address{"SomeWhere 12", "12312, Greece"}

	// Преобразовать структуру в XML
	output, err := xml.MarshalIndent(r, " ", " ")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Преобразовать XML в срез байтов, предварительно добавив заголовок
	output = []byte(xml.Header + string(output))
	
	// Записать в поток вывода срез байтов
	_, _ = os.Stdout.Write(output)
	_, _ = os.Stdout.Write([]byte("\n"))
}
