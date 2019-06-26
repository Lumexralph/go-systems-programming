package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
)

var (
	MAX      = 90
	MIN      = 0
	seedSize = 10
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	var seed int64

	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Printf("usage: %s length \n", filepath.Base(arguments[0]))
		os.Exit(1)
	}

	LENGTH, _ := strconv.ParseInt(arguments[1], 10, 64)
	f, err := os.Open("/dev/random")
	if err != nil {
		log.Fatalf("Cannot open the file %s: %s", arguments[1], err)
	}

	// build an int64 instead of a series of bytes.
	// This is an example of having to read from a binary file into Go types.
	binary.Read(f, binary.LittleEndian, &seed)
	rand.Seed(seed)
	f.Close()

	fmt.Println("Seed: ", seed)

	startChar := "!"
	var i int64

	for i = 0; i < LENGTH; i++ {
		anInt := int(random(MIN, MAX))

		// convert an integer number into an ASCII character
		newChar := string(startChar[0] + byte(anInt))
		if newChar == " " {
			i = i - i
			continue
		}
		fmt.Print(newChar)
	}
	fmt.Println()

}
