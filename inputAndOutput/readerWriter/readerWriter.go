package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Please provide a filename")
		os.Exit(1)
	}

	filename := arguments[1]

	// get information about a file if it exists
	_, err := os.Stat(filename)
	if err != nil {
		fmt.Printf("Error on file: %s", err)
		os.Exit(1)
	}

	// if file exists, attempt to open it
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Cannot open file: %s", err)
		os.Exit(-1)
	}
	defer f.Close()

	chars := countChars(f)
	filename = filename + ".Count"

	// attempt to create this file
	f, err = os.Create(filename)
	if err != nil {
		fmt.Printf("Creating file error: %s", err)
		os.Exit(1)
	}
	defer f.Close()

	writeNumberOfChars(f, chars)
}

func countChars(r io.Reader) int {
	buf := make([]byte, 16)
	total := 0

	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return 0
		}

		// if it gets to the end of the file
		if err == io.EOF {
			break
		}
		total = total + n
	}

	return total
}

func writeNumberOfChars(w io.Writer, num int) {
	fmt.Fprintf(w, "%d \n", num)
}
