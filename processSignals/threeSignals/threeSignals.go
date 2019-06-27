package main

import (
	"time"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func handleSignal(signal os.Signal) {
	fmt.Println("Got ", signal)
}

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	go func() {
		for {
			sig := <-sigs

			switch sig {
			case syscall.SIGINT:
				handleSignal(sig)
			case syscall.SIGTERM:
				handleSignal(sig)
			case syscall.SIGHUP:
				fmt.Println("Got: ", sig)
				os.Exit(-1)
			}
		}
	}()

	for {
		fmt.Println(".")
		time.Sleep(10 * time.Second)
	}
}
