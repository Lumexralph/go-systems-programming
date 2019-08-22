package main

import (
	"fmt"
	"io"
)

func readSocket(r io.Reader) {

	// create a memory buffer where to read data
	// from the socket file and store into the buffer
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf[:])
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("--> Client %v \n", string(buf[0:n]))
	}
}
