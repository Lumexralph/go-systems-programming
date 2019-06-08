package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()
	flags := flag.Args()

	if len(flags) == 0 {
		fmt.Println("Not enough arguments!")
		os.Exit(1)
	}

	path := flags[0]

	err := filepath.Walk(path, walkFunction)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func walkFunction(path string, info os.FileInfo, err error) error {
	// check if the file/directory exists
	// and get the info about it
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	mode := fileInfo.Mode()
	if mode.IsDir() || mode.IsRegular() {
		fmt.Println(path)
	}

	return nil // there was no error
}
