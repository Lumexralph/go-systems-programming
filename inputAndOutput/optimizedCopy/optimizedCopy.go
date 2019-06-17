package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

var BUFFERSIZE int64

func main() {
	arguments := os.Args
	if len(arguments) != 4 {
		fmt.Printf("usage: %s source destination BUFFERSIZE \n",
			filepath.Base(arguments[0]))
		os.Exit(1)
	}

	source := arguments[1]
	destination := arguments[2]
	BUFFERSIZE, err := strconv.ParseInt(arguments[3], 10, 64)

	fmt.Printf("Copying source %s to destination %s \n", source, destination)
	err = copy(source, destination, BUFFERSIZE)
	if err != nil {
		fmt.Printf("File copying failed: %q \n", err)
	}
}

func copy(src, dst string, BUFFERSIZE int64) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	_, err = os.Stat(dst)
	if err == nil {
		return fmt.Errorf("File %s already exists", dst)
	}

	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	if err != nil {
		panic(err)
	}

	buf := make([]byte, BUFFERSIZE) // create a  buffer

	for {
		// read the content of the file sequentially by bufferring
		n, err := sourceFile.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}

		// if the reading gets to the end of the file io.EOF is
		// returned but it doesn't cause or lead to an error
		if n == 0 {
			break
		}

		if _, err := destinationFile.Write(buf[:n]); err != nil {
			return err
		}
	}
	return err
}
