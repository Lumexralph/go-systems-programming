package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		log.Fatalln("Please provide a port number")
	}

	PORT := ":" + arguments[1]

	UDPAddr, err := net.ResolveUDPAddr("udp", PORT)
	if err != nil {
		log.Fatalln(err)
	}

	// make the app listening to connection on a port
	UDPConnection, err := net.ListenUDP("udp", UDPAddr)
	if err != nil {
		log.Fatalln(err)
	}

	defer UDPConnection.Close()

	// create a buffer to temporary store the data
	buffer := make([]byte, 1024)

	// keep the server on to continuosly listen
	for {
		n, UDPAddr, err := UDPConnection.ReadFromUDP(buffer)
		fmt.Printf("--> %s", string(buffer[0:n]))
		data := []byte(buffer[0:n])

		// write to the connection to create a response
		n, err = UDPConnection.WriteToUDP(data, UDPAddr)
		if err != nil {
			log.Fatalln(err)
		}

		// UDP connection does not need to call a function similar to net.Accept() in TCP
		if strings.TrimSpace(string(data)) == "STOP" {
			fmt.Println("EXiting the UDP Server.....")
			return
		}
	}
}
