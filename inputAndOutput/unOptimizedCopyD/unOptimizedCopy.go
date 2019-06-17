package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) != 3 {
		fmt.Println("Please provide 2 arguments for destination and source file")
		os.Exit(1)
	}

	destinationFile := arguments[2]
	sourceFile := arguments[1]

	nBytes, err := copy(sourceFile, destinationFile)
	if err != nil {
		fmt.Printf("The copy operation failed, %q \n", err)
		os.Exit(1)
	} else {
		fmt.Printf("Copied %d bytes to destination file \n", nBytes)
	}

}

func copy(src, dst string) (int64, error) {
	// get information about the source if it exists
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	// check if the source is not a regular file
	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	sourceFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destinationFile.Close()

	nBytes, err := io.Copy(destinationFile, sourceFile)

	return nBytes, err
}
