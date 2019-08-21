package main

import (
	"net"
	"fmt"
	"os"
)

// This is where functions that servers incoming request
// is implemented
func main() {

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a socket file")

		// make a sys call to exit the process
		os.Exit(100)
	}

	socketFile := arguments[1]

	listener, err := net.Listen("unix", socketFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(100)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			os.Exit(100)
		}

		// run the server in another goroutine, thread or
		// in a way process but child process
		// because the server.go and main.go belong to same
		// package, we can access the function without importing it
		go echoServer(conn)
	}

}
