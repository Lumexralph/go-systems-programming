package main

import (
	"fmt"
	"net"
)

func echoServer(c net.Conn) {

	// need an infinite loop to keep the server
	// up and always listening
	for {

		// get 1024 bytes from memory
		buf := make([]byte, 1024)

		// read from the stream of bytes
		nr, err := c.Read(buf)
		if err != nil {

			// if error comes up, quit the server
			return
		}

		// read the data from the beginning of the base
		// addres towards the length of thr byte read
		data := buf[0:nr]
		fmt.Printf("-->: %v \n", string(data))

		// write from the stream of bytes
		_, err = c.Write(data)
		if err != nil {
			fmt.Println(err)
		}
	}
}
