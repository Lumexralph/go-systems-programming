/*This section will present a Go program that converts tab characters
to space characters in files and vice versa! This is the job that is
usually done by a text editor*/
package main

import (
	"strings"
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args

	if len(arguments) != 3 {
		fmt.Printf("Usage: %s [-t|-s] filename \n", filepath.Base(arguments[0]))
		os.Exit(1)
	}

	convertTabs := false
	convertSpaces := false
	newLine := ""

	option := arguments[1]
	filename := arguments[2]

	if option == "-t" {
		convertTabs = true
	} else if option == "-s" {
		convertSpaces = true
	} else {
		fmt.Println("Unknown option!")
		os.Exit(1)
	}

	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file %s: %s", filename, err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("Error reading file %s", err)
			os.Exit(1)
		}

		if convertTabs {
			newLine = strings.Replace(line, "\t", "    ", -1)
		} else if convertSpaces {
			newLine = strings.Replace(line, "    ", "\t", -1)
		}

		fmt.Print(newLine)
	}
}
