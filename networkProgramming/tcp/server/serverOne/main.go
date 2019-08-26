package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	arguments := os.Args
	if len(arguments) == 1 {
		log.Fatalln("Please provide port number")
	}

	PORT := ":" + arguments[1]

	// make the serve listen on the address
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalln(err)
	}

	// ensure to close the connection after executing the function
	defer listener.Close()

	// only after a successful call to accept can the
	// tcp client make connection to the server
	conn, err := listener.Accept()
	if err != nil {
		log.Fatalln(err)
	}

	for {
		// read data from the connection stream
		// read data from the client
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("--> TCP Server %v \n", string(data))

		// return a response to the client through a
		// write to the connection stream between the server and client
		conn.Write([]byte(data))
		if strings.TrimSpace(string(data)) == "STOP" {
			fmt.Println("Exiting the TCP server")
			return
		}

	}
}
