package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var readValue = make(chan int)
var writeValue = make(chan int)

func setValue(newValue int) {
	writeValue <- newValue
}

func getValue() int {
	return <-readValue
}

func monitor() {
	var value int

	for {
		select {
		case newValue := <-writeValue:
			value = newValue
			fmt.Printf("%d : ", value)
		case readValue <- value:

		}
	}
}

func main() {
	var waitGroup sync.WaitGroup

	rand.Seed(time.Now().Unix())
	go monitor()

	for r := 0; r < 20; r++ {
		waitGroup.Add(1)

		go func() {
			defer waitGroup.Done()

			setValue(rand.Intn(100))
		}()
	}

	waitGroup.Wait()
	fmt.Printf("\n Last Value: %d \n", getValue())
}
