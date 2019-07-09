package main

import (
	"fmt"
	"sync"
	"time"
)

type secret struct {
	sync.RWMutex
	counter  int
	password string
}

var password = secret{
	counter:  1,
	password: "gamer",
}

func change(c *secret, pass string) {
	c.Lock()

	fmt.Println("LChange")
	time.Sleep(20 * time.Second)
	c.counter = c.counter + 1
	c.password = pass

	c.Unlock()
}

func show(c *secret) string {
	fmt.Println("LShow")
	time.Sleep(time.Second)

	c.RLock()
	defer c.RUnlock()
	return c.password
}

func counts(c *secret) int {
	c.RLock()
	defer c.RUnlock()
	return c.counter
}

func main() {
	fmt.Println("Password: ", show(&password))

	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println("Go Password: ", show(&password))
		}()
	}

	go func() {
		change(&password, "123456")
	}()

	fmt.Println("Password: ", show(&password))
	time.Sleep(time.Second)
	fmt.Println("Counter: ", counts(&password))
}
