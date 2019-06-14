package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Please provide a filename")
		os.Exit(1)
	}

	filename := arguments[1]

	byteSlice := []byte("Go systems programming \n")
	ioutil.WriteFile(filename, byteSlice, 0644) // write data to the file

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	anotherByteSlice := make([]byte, 100)
	n, err := f.Read(anotherByteSlice)
	fmt.Printf("Read %d bytes from the string %s \n", n, anotherByteSlice)
}
