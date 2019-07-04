package main

import (
	"fmt"
	"strconv"
	"time"
)

func firstResponder(out chan string) {
	replyToOtherResponder := "How are you: Ping!"

	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		out <- replyToOtherResponder + strconv.Itoa(1)
	}
}

func secondResponder(out chan string) {
	replyToOtherResponder := "I am fine: Pong!"

	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Second)
		out <- replyToOtherResponder +  strconv.Itoa(1)
	}
}

func main() {
	medium1 := make(chan string)

	go firstResponder(medium1)
	go secondResponder(medium1)

	for i := 0; i < 5; i++ {
		fmt.Println("I am from first responder: ", <-medium1)
		time.Sleep(1 * time.Second)
		fmt.Println("I am from second responder: ", <-medium1)
		time.Sleep(1 * time.Second)
	}

	close(medium1)

}
