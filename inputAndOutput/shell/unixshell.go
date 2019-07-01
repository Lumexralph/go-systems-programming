package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const VERSION = "0.2"

func main() {
	// present prompt
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")

	for scanner.Scan() {
		// Read a line
		line := scanner.Text()
		words := strings.Split(line, " ")

		// Get the first word of the line
		command := words[0]
		switch command {
		case "exit":
			fmt.Println("Exiting .......")
			os.Exit(0)
		case "version":
			fmt.Println(VERSION)
		default:
			fmt.Println(line)
		}

		fmt.Print("> ")
	}

	

	

	// if it is a built-in shell, execute the command

	// otherwise, echo the command
}
