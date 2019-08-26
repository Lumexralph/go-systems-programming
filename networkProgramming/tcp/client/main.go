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
		log.Fatalln("Please provide the host:port")
	}

	address := arguments[1]

	// try to connect to the address
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		// create a reader interface to read data from
		reader := bufio.NewReader(os.Stdin)

		fmt.Print(">> ")

		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		// send data to the connection or write to it
		fmt.Fprintf(conn, text + " \n")

		// get data from the server and read from the writer
		// which is the connection between the client and server
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("--> client: %s", message)
		if strings.TrimSpace(text) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}

	}
}
