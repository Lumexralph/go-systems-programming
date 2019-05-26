package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {

	var s [3]string
	s[0] = "1 b 3 1 a a b"
	s[1] = "11 g 13 h Babe 16 a a"
	s[2] = "-1 b -3 -4 a -6"

	// Create the regex pattern to be matched
	parse, err := regexp.Compile("[bB]")

	if err != nil {
		fmt.Printf("Error creating the RegEx pattern: %s\n", err)
		os.Exit(-1)
	}

	for i := 0; i < len(s); i++ {
		temp := parse.ReplaceAllString(s[i]," Ralph")
		fmt.Println(temp)
	}
}
