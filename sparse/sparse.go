package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	arguments := os.Args

	if len(arguments) != 3 {
		fmt.Printf("usage: %s SIZE filename \n", filepath.Base(arguments[0]))
		os.Exit(1)
	}

	SIZE, err := strconv.ParseInt(arguments[1], 10, 64)
	filename := arguments[2]

	_, err = os.Stat(filename)
	if err == nil {
		fmt.Printf("File %s already exists. \n", filename)
		os.Exit(1)
	}

	fd, err := os.Create(filename)
	if err != nil {
		log.Fatal("Failed to create file")
	}

	_, err = fd.Seek(SIZE - 1, 0)
	if err != nil {
		fmt.Println("err")
		log.Fatal("Failed to seek on file")
	}

	_, err = fd.Write([]byte{0})
	if err != nil {
		fmt.Println("err")
		log.Fatal("Write operation failed")
	}
	err = fd.Close()
	if err != nil {
		fmt.Println("err")
		log.Fatal("Failed to close file")
	}

}
