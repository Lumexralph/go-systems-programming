package main

import (
	"fmt"
	"time"
)

func writeChannel(c chan<- int, x int) {
	fmt.Println(x)

	// the data will be lost because currently nobody
	// reads the channel in the program even though data
	// is written or sent to it
	c <- x
	close(c)
	fmt.Println(x)
}

func main() {
	c := make(chan int)

	go writeChannel(c, 7)
	time.Sleep(2 * time.Second)
	val := <-c
	fmt.Println("Read: ", val)
	time.Sleep(2 * time.Second)

	_, ok := <-c
	if ok {
		fmt.Println("Channel is still open!")
	} else {
		fmt.Println("Channel is Closed!")
	}
}
