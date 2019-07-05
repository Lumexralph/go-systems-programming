package main

import (
	"fmt"
)

func main() {
	numbers := make(chan int, 5)
	counter := 10

	// labelling the loop so as to stop the loop once there is
	// no more data in the channel because using ordinary break
	// will beak the case/select but not the loop
firstLoop:
	for i := 0; i < counter; i++ {
		select {
		case numbers <- i:
		default:
			fmt.Printf("Not enough space for the value - %d - 10 \n", i)
			break firstLoop
		}

	}

nextLoop:
	for i := 0; i < counter*2; i++ {
		select {
		case num := <-numbers:
			fmt.Println(num)
		default:
			fmt.Println("No data left to be received from numbers channel")
			break nextLoop
		}
	}
}
