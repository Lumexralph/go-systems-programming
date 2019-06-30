package main

import (
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("Lumexky", "I am Lumexky")
	fmt.Println(match)

	match, _ = regexp.MatchString("lumexky", "I am Lumexky in the room")
	fmt.Println(match)

	// read and parse the regex pattern
	// it will create a regex pattern if successful
	// this can be used to parse other string or text
	// afterwards
	parse, err := regexp.Compile("[Ll]umexky")
	if err != nil {
		fmt.Printf("Error compiling RegExp: %s\n", err)
	} else {
		fmt.Println(parse.MatchString("Ogundele Olumide lumexky"))
		fmt.Println(parse.MatchString("Ogundele Olumide Lumexky"))
		fmt.Println(parse.MatchString("Ogundele Olumide l umexky"))
		fmt.Println(parse.ReplaceAllString("Ogundele Olumide Lumexky lumexky", "Ralph"))
	}
}
