package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	sum := 0

	// ignore the first element in the arguments array
	// which is the executable file
	for i := 1; i < len(arguments); i++ {
		temp, err := strconv.Atoi(arguments[i])

		if err == nil {
			sum = sum + temp
		} else {
			fmt.Println("Ignoring the argument...")
		}

	}
	fmt.Println("The Sum is: ", sum)
}
