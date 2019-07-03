package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Printf("Waiting for goroutines....")

	var waitGroup sync.WaitGroup
	// the number of goroutine to wait for
	waitGroup.Add(5)      // will cause an error just to proof that there
							   // should be an equivalent goroutines to call waitGroup.Done()
							   // no more or less but equivalent else you get error

	var i int64

	for i = 0; i < 10; i++ {

		go func (x int64)  {
			defer waitGroup.Done()
			fmt.Printf("%d ", x)
		}(i)
	}

	waitGroup.Wait()
	fmt.Println("\n Exiting......")

}
