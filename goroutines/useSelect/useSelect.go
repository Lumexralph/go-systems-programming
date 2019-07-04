package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func createNumber(max int, randomNumberChannel chan<- int, finishedChannel chan bool) {
	for {
		select {
		case randomNumberChannel <- rand.Intn(max):
		case x := <-finishedChannel:
			if x {
				close(randomNumberChannel)
				close(finishedChannel)
				return
			}
		}
	}
}

func main() {
	arguments := os.Args

	rand.Seed(time.Now().Unix())
	randomNumberChannel := make(chan int)
	finishedChannel := make(chan bool)

	if len(arguments) != 3 {
		fmt.Printf("usage: %s count max \n", filepath.Base(arguments[0]))
		os.Exit(1)
	}

	n1, _ := strconv.ParseInt(arguments[1], 10, 64)
	n2, _ := strconv.ParseInt(arguments[2], 10, 64)
	count := int(n1)
	max := int(n2)

	fmt.Printf("Going to create %d random numbers. \n", count)

	go createNumber(max, randomNumberChannel, finishedChannel)

	for i := 0; i < count; i++ {
		fmt.Printf("%d ", <-randomNumberChannel)
	}

	finishedChannel <- false
	_, ok := <-randomNumberChannel
	if ok {
		fmt.Println("Channel is open!...")
	} else {
		fmt.Println("Channel is closed!...")
	}

	finishedChannel <- true
	_, ok = <-randomNumberChannel
	if ok {
		fmt.Println("Channel is open!...")
	} else {
		fmt.Println("Channel is closed...")
	}
}
