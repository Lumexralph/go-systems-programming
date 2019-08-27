package main

import (
	"log"
	"net"
	"net/rpc"

	"github.com/Lumexralph/go_systems_programming/networkProgramming/rpc/sharedRPC"
)

type MyInterface int

// Add is the implementation of the first function that
// will be offered to the RPC clients—you can have as many
// functions as you want, provided that they are included
// in the interface.
func (t *MyInterface) Add(arguments *sharedRPC.MyInts, reply *int) error {

	s1 := 1
	s2 := 1

	if arguments.S1 == true {
		s1 = -1
	}

	*reply = s1*int(arguments.A1) + s2*int(arguments.A2)

	return nil
}

// Subtract - is the second function that is offered to the RPC clients by this RPC server
func (t *MyInterface) Subtract(arguments *sharedRPC.MyInts, reply *int) error {

	s1 := 1
	s2 := 2

	if arguments.S1 == true {
		s1 = -1
	}

	if arguments.S2 == true {
		s2 = -1
	}

	*reply = s1*int(arguments.A1) - s2*int(arguments.A2)

	return nil
}

func main() {

	PORT := ":1234"

	// create a pointer(that holds the address) to store
	// the interface type in memory
	myInterface := new(MyInterface)

	// register the interface that will receive RPC function call
	// in order to be able to serve the desired interface.
	rpc.Register(myInterface)

	// As our RPC server uses TCP, you need to make calls
	// to net.ResolveTCPAddr() and net.ListenTCP()
	TCPAddr, err := net.ResolveTCPAddr("tcp", PORT)
	if err != nil {
		log.Fatalln(err)
	}

	// listen and bind the server to the address
	listener, err := net.ListenTCP("tcp", TCPAddr)

	// I am goofing around here
	type errors []error
	a := make(errors, 8)

	if err != nil {
		a[0] = err
		log.Fatalln(a)
	}

	for {
		connection, err := listener.Accept()
		if err != nil {
			continue
		}

		rpc.ServeConn(connection)
	}

}
