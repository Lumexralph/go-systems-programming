// save data in the CSV format
package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Need just one filename")
		os.Exit(-1)
	}

	filename := arguments[1]

	// get the information of the file if it exists
	_, err := os.Stat(filename)
	if err == nil {
		fmt.Printf("File %s already exist", filename)
		os.Exit(1)
	}

	output, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer output.Close()

	// using a 2-dimensional slice
	inputData := [][]string{
		{"M", "T", "I."},
		{"D", "T", "I."},
		{"M", "T", "D."},
		{"V", "T", "D."},
		{"A", "T", "D."},
	}

	// create writer interface for csv using the file descriptor (INode)
	writer := csv.NewWriter(output)

	for _, record := range inputData {
		err := writer.Write(record)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	writer.Flush()

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1
	allRecords, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, row := range allRecords {
		fmt.Printf("%s:%s:%s \n", row[0], row[1], row[2])
	}
}
