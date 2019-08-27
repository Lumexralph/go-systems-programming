package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"

	"github.com/Lumexralph/go_systems_programming/networkProgramming/rpc/sharedRPC"
)

func main() {

	arguments := os.Args
	if len(arguments) == 1 {
		log.Fatalln("Please provide a host:port arguments")
	}

	address := arguments[1]

	// make a request to create an RPC connection
	//  tp the RPC server
	connection, err := rpc.Dial("tcp", address)
	if err != nil {
		log.Fatalln(err)
	}

	args := sharedRPC.MyInts{
		18,
		20,
		true,
		false,
	}

	var reply int

	// make a function call from the shared interface
	//  you use the Call() function to execute the desired
	// function in the RPC server. The result of the
	// MyInterface.Add() function is stored in the reply
	// variable, which was previously declared.
	err = connection.Call("MyInterface.Add", args, &reply)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Reply (Add): %d \n", reply)

	err = connection.Call("MyInterface.Subtract", args, &reply)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Reply (Subtract): %d \n", reply)
}
