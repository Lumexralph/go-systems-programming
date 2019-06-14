package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	arguments := os.Args

	if len(arguments) != 2 {
		fmt.Println("Please provide a filename")
		os.Exit(1)
	}

	filename := arguments[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening opening file: %s", err)
		os.Exit(1)
	}
	defer file.Close()
	
	// create a scanner that will run through the file on demand
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if scanner.Err() != nil {
			fmt.Printf("Error occurred reading the file: %s", err)
			os.Exit(1)
		}

		fmt.Println(line)
	}
}
