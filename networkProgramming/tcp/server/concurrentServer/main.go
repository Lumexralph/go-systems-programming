package main

import (
	"log"
	"net"
	"os"
)

/*
each client connection will be assigned to a new
 goroutine that will serve the client request.
 Note that although TCP clients initially connect to the
 same port, they are served using a different port number
 than the main port number of the server—this is
 automatically handled by TCP and is the way TCP works.
*/

func main() {

	arguments := os.Args
	if len(arguments) == 1 {
		log.Fatalln("Please provide a port number")
	}

	PORT := ":" + arguments[1]

	// listen for networkconnection request
	// or rather bind the port to the server
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalln(err)
	}

	defer listener.Close()

	for {
		// make the listener accept connection request
		// using the accept method of the Listener Interface
		connection, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		// handle multiple requests from clients using
		// other child processes by using goroutines
		go handleConnection(connection)
	}

}
