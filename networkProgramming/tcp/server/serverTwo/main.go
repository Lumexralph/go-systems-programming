package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	arguments := os.Args
	if len(arguments) == 1 {
		log.Fatalln("Please provide a port number")
	}

	address := "localhost" + ":" + arguments[1]
	TCPAddress, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		log.Fatalln(err)
	}

	// make a listen to any connection to the address
	listener, err := net.ListenTCP("tcp", TCPAddress)
	if err != nil {
		log.Fatalln(err)
	}

	// make a buffer memory to hold data
	buffer := make([]byte, 1024)

	// continuously open the connection for the server
	for {

		TCPConn, err := listener.AcceptTCP()
		if err != nil {
			log.Fatalln(err)
		}

		// read data from the connection stream / writer to the buffer
		n, err := TCPConn.Read(buffer)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf(">> %s", string(buffer[0:n]))

		// make a response to the connection / write to the writer
		n, err = TCPConn.Write(buffer)
		// close the connection to free up resources
		TCPConn.Close()

		if err != nil {
			log.Fatalln(err)
		}
	}
}
