package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	// the first element will always be the file or program name
	if len(arguments) == 1 {
		fmt.Println("Please provide an argument!")
		os.Exit(1)
	}

	file := arguments[1]
	// check if the file exists in the filesystem
	fileInfo, err := os.Stat(file)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	mode := fileInfo.Mode()
	fmt.Printf("%s : %s  %d \n", file, mode, mode)

}
