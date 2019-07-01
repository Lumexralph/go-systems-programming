package main

import (
	"fmt"
	"io"
	"net"
	"strconv"
	"time"
)

func readSocket(r io.Reader) {
	buf := make([]byte, 1024)

	for {
		n, _ := r.Read(buf[:])
		fmt.Print("Read: ", string(buf[0:n]))
	}
}

func main() {
	conn, _ := net.Dial("unix", "/tmp/aSocket.sock")
	defer conn.Close()

	go readSocket(conn)

	n := 0
	for {
		message := []byte("Hi there user: " + strconv.Itoa(n) + "\n")
		_, _ = conn.Write(message)

		// delay the process
		time.Sleep(5 * time.Second)
		n = n + 1
	}
}
