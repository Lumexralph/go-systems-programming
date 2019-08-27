package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func handleConnection(c net.Conn) {

	for {
		// create a reader from the connection stream to
		// be able to the read the data in the stream
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("-> %s", string(netData))

		// write back to the connection stream to create a response
		c.Write([]byte(netData))

		if strings.TrimSpace(string(netData)) == "STOP" {
			break
		}
	}

	// delay the process for 3 seconds
	// The time delay at the end of it is used for giving
	// you the necessary time to connect with another TCP client
	time.Sleep(3 * time.Second)

	// close the connection to free up resources
	c.Close()
}
