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
		log.Fatalln("Please provide a domain")
	}

	domain := arguments[1]

	NS, err := net.LookupNS(domain)
	if err != nil {
		log.Fatalln(err)
	}

	for _, NS := range NS {
		fmt.Println(NS.Host)
	}
}
