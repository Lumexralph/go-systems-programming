package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Please provide an integer !")
		os.Exit(1)
	}
	number, err := strconv.ParseInt(arguments[1], 10, 64)

	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.LittleEndian, number)
	if err != nil {
		fmt.Println("Little Endian: ", err)
	}

	fmt.Printf("%d is %x in Little Endian \n", number, buf)
	// clean up the the buffer writer
	buf.Reset()

	err = binary.Write(buf,binary.BigEndian, number)
	if err != nil {
		fmt.Println("Big Endian error: ", err)
	}

	fmt.Printf("And %x in Big Endian \n", buf)

}
