package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type RecordCsv struct {
	Name       string
	Surname    string
	Number     string
	LastAccess string
}

var data []RecordCsv

func readCSVFile(filepath string) ([][]string, error) {
	_, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// CSV-файл читается весь сразу
	// тип данных lines - [][]string
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, nil
	}

	return lines, nil
}

func saveCSVFile(filepath string) error {
	csvFile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	// изменение разделителя по умолчанию на табуляцию
	csvWriter.Comma = '\t'
	for _, row := range data {
		temp := []string{
			row.Name,
			row.Surname,
			row.Number,
			row.LastAccess,
		}
		_ = csvWriter.Write(temp)
	}
	csvWriter.Flush()
	return nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("csvData input output!")
		return
	}

	input := os.Args[1]
	output := os.Args[2]
	lines, err := readCSVFile(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	// данные CSV считываются по столбцам - каждая строка
	// представляет собой срез
	for _, line := range lines {
		temp := RecordCsv{
			Name:       line[0],
			Surname:    line[1],
			Number:     line[2],
			LastAccess: line[3],
		}
		data = append(data, temp)
		fmt.Println(temp)
	}

	err = saveCSVFile(output)
	if err != nil {
		fmt.Println(err)
		return
	}
}
