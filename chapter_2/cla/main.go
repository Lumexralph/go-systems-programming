// package that prints all its command-line arguments
// including the name of the executable file

package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args

	for i := 0; i < len(arguments); i++ {
		fmt.Println(arguments[i])
	}
}
