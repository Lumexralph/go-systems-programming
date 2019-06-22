package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	minusCOL := flag.Int("COL", 1, "Column")
	flag.Parse()

	flags := flag.Args()

	if len(flags) == 0 {
		fmt.Printf("usage: %s <file1> [<file2> [...<fileN>]] \n",
			filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	column := *minusCOL
	if column < 0 {
		fmt.Println("Invalid Column number!")
		os.Exit(1)
	}

	IPs := make(map[string]int)
	for _, filename := range flags {
		fmt.Println("\t\t", filename)
		f, err := os.Open(filename)
		if err != nil {
			fmt.Printf("Error opening file: %s", err)
			continue
		}
		defer f.Close()

		// create a buffer reader
		r := bufio.NewReader(f)
		for {
			line, err := r.ReadString('\n')

			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("Error reading file: %s", err)
				continue
			}

			data := strings.Fields(line)
			ip := data[column-1]
			trial := net.ParseIP(ip)
			if trial.To4() == nil {
				continue
			}

			_, ok := IPs[ip]
			if ok {
				IPs[ip] = IPs[ip] + 1
			} else {
				IPs[ip] = 1
			}
		}

	}

	for key, value := range IPs {
		fmt.Printf("%s - %d \n", key, value)
	}
}
