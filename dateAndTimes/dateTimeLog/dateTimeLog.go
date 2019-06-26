package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("PLease provide one log file to process")
		os.Exit(-1)
	}

	filename := flag.Arg(0)
	f, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file %s", err)
		os.Exit(1)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	regEX := regexp.MustCompile(`.*\[(\d\d\/\w+/\d\d\d\d:\d\d:\d\d:\d\d.*)\] .*`)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("Error reading file %s \n", err)
		}

		if regEX.MatchString(line) {
			match := regEX.FindStringSubmatch(line)
			// once you have a match, parse the date and time you found
			d1, err := time.Parse("02/Jan/2006:15:04:05 -0700", match[1])
			if err != nil {
				fmt.Println(err)
			}

			// convert it to the desired format
			newFormat := d1.Format(time.RFC3339)
			fmt.Print(strings.Replace(line, match[1], newFormat, 1))
		}
	}
}
