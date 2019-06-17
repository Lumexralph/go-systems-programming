package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) != 3 {
		fmt.Println("Please provide two command line arguments")
		os.Exit(1)
	}

	destinationFile := arguments[1]
	sourceFile := arguments[2]

	input, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ioutil.WriteFile(destinationFile, input, 0644)
	if err != nil {
		fmt.Printf("Error creating newfile %s: %q", destinationFile, err)
		os.Exit(1)
	}
}
