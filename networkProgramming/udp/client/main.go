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
		log.Fatalln("Please provide a host:port argument")
	}

	addr := arguments[1]

	UDPAddr, err := net.ResolveUDPAddr("udp", addr)

	// open the connection stream to the server
	UDPConnection, err := net.DialUDP("udp", nil, UDPAddr)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("The UDP Server is %s \n", UDPConnection.RemoteAddr().String())

	// close the connection to free up resources when done
	defer UDPConnection.Close()

	data := []byte("Hello UDP Echo Server! \n")

	// write/send data to the connection stream and make a response to the server
	_, err = UDPConnection.Write(data)
	if err != nil {
		log.Fatalln(err)
	}

	// create a buffer, some section in memory to temporarily
	// store the data read from the server or the connection stream
	buffer := make([]byte, 1024)
	n, _, err := UDPConnection.ReadFromUDP(buffer)

	fmt.Printf("Reply: %s", string(buffer[:n]))


}
