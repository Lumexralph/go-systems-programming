package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// get the command line arguments
	arguments := os.Args
	minusI := false

	for i := 0; i < len(arguments); i++ {
		if strings.Compare(arguments[i], "-i") == 0 {
			minusI = true
			break
		}
	}

	if minusI {
		fmt.Println("Got the -i parameter")
		fmt.Print("y/n: ")
		var answer string
		// read the user input from the input stream
		// Since Go passes its variables by value, we have
		// to use a pointer reference here in order to save
		//  the user input
		fmt.Scanln(&answer)
		fmt.Println("You entered:  ", answer)
	} else {
		fmt.Println("The -i parameter is not set")
	}

}
