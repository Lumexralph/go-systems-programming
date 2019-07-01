package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func catFile(filename string) error {
	fileHandler, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileHandler.Close()

	scanner := bufio.NewScanner(fileHandler)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return nil
}

func main() {
	filename := ""
	arguments := os.Args
	if len(arguments) == 1 {
		io.Copy(os.Stdout, os.Stdin)
		os.Exit(0)
	}

	filename = arguments[1]
	err := catFile(filename)
	if err != nil {
		fmt.Println(err)
	}
}
