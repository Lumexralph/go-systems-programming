/*

Main package will read inputs from the user
keep track of the minimum and maximum value
When the inputted value is zero, it ends the program
and output the minimum and maximum value

*/
package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// start the program
	min, max := receiveInput()
	fmt.Printf("Min: %d, Max: %d \n", min, max)
}

// receiveInput - will read input from user
// keep track of min and max value
// end execution when the user inputs 0
func receiveInput() (min, max int) {
	// a loop to make it keep requesting for input
	var input string
	for {
		fmt.Print("Please provide an integer: ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			log.Fatalln("Sorry, error occured while reading your input", err)
		}
		// convert to an integer
		value, err := strconv.Atoi(input)
		if err != nil {
			log.Println("Please, provide an integer")
			continue
		}

		switch {
		case value == 0:
			return
		case value < min:
			min = value
		case max < value:
			max = value
		}
	}
}
