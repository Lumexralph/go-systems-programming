package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
)

type file struct {
	filename   string
	lines      int
	words      int
	characters int
	err        error
}

func monitor(values <-chan file, count int) {
	var (
		totalWords = 0
		totalLines = 0
		totalChars = 0
	)

	for i := 0; i < count; i++ {
		x := <-values
		totalWords = totalWords + x.words
		totalLines = totalLines + x.lines
		totalChars = totalChars + x.characters
		if x.err == nil {
			fmt.Printf("\t %d \t", x.lines)
			fmt.Printf(" %d \t", x.words)
			fmt.Printf(" %d \t", x.lines)
			fmt.Printf(" %s \n", x.filename)
		} else {
			fmt.Printf("\t %s \n", x.err)
		}
	}

	fmt.Printf("\t %d \t", totalLines)
	fmt.Printf(" %d \t", totalWords)
	fmt.Printf(" %d \t total \n", totalChars)
}

func count(filename string, out chan<- file) {
	var (
		err    error
		nLines = 0
		nChars = 0
		nWords = 0
	)

	f, err := os.Open(filename)
	defer f.Close()

	if err != nil {
		newFileValue := file{
			filename:   filename,
			lines:      0,
			words:      0,
			characters: 0,
			err:        err,
		}

		out <- newFileValue
	}

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("Error reading file %s \n", err)
		}

		nLines++
		r := regexp.MustCompile("[^\\s]+")
		for range r.FindAllString(line, -1) {
			nWords++
		}
		nChars += len(line)
	}

	newFileValue := file{
		filename:   filename,
		lines:      nLines,
		words:      nWords,
		characters: nChars,
		err:        nil,
	}

	out <- newFileValue
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Printf("usage: %s <file1> [<file2> [... <fileN]]] \n",
			filepath.Base(arguments[0]))
		os.Exit(1)
	}

	values := make(chan file, len(arguments[1:]))
	for _, filename := range arguments[1:] {
		go func(filename string) {
			count(filename, values)
		}(filename)
	}

	monitor(values, len(arguments[1:]))
}
