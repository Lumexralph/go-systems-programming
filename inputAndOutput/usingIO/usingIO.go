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
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening file %s: %s \n", filename, err)
		os.Exit(1)
	}
	defer f.Close()

	// make a buffer to read 8 bytes of data
	buf := make([]byte, 8)
	if _, err = io.ReadFull(f, buf); err != nil {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}

	io.WriteString(os.Stdout, string(buf))
	fmt.Println()
}
