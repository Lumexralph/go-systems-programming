package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a socket file")
		os.Exit(100)
	}

	socketFile := arguments[1]

	// make a request to the socket server
	conn, err := net.Dial("unix", socketFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(100)
	}

	go readSocket(conn)

	for {
		_, err = conn.Write([]byte("Sending Request to Unix Server"))
		if err != nil {
			fmt.Println(err)
			os.Exit(100)
		}

		time.Sleep(1 * time.Second)
	}
}
