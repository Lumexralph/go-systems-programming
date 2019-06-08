package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// printDir := flag.Bool("d", false, "print directory")
	// printFile := flag.Bool("f", false, "print file")

	// flag.Parse()

	// flags := flag.Args()

	// if len(flags) == 0 {
	// 	fmt.Println("Please provide an argument")
	// 	os.Exit(1)
	// }

	// path := flags[0]
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Not enough arguments!")
		os.Exit(1)
	}

	path := arguments[1]
	err := filepath.Walk(path, walkFunction)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func walkFunction(path string, info os.FileInfo, err error) error {
	// check if the file/directory exists
	// and get the info about it
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	mode := fileInfo.Mode()

	switch {
	case mode.IsDir():
		fmt.Println(path)
	case mode.IsRegular():
		fmt.Println("**********")
		// default:
		// 	fmt.Println(path)
	}

	return nil // there was no error
}
