package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filename := ""
	var f *os.File
	arguments := os.Args
	if len(arguments) == 1 {
		f = os.Stdin
	} else {
		filename = arguments[1]
		fileHandler, err := os.Open(filename)
		if err != nil {
			fmt.Printf("Error opening file %s: %s", filename, err)
			os.Exit(1)
		}
		f = fileHandler
	}
	defer f.Close()
	// can be further exploited to build a REPL for any language
	fmt.Print("> ")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		fmt.Print("> ")
	}
}
