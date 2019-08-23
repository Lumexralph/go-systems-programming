package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	// get the command line arguments
	arguments := os.Args
	if len(arguments) == 1 {
		log.Fatalln("Please provide a URL to lookup")
	}

	hostname := arguments[1]
	IPs, err := net.LookupHost(hostname)
	if err != nil {
		log.Fatalln(err)
	}

	for _, IP := range IPs {
		fmt.Println(IP)
	}
}
