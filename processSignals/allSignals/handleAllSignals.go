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
	signal.Notify(sigs)

	go func() {
		for {
			sig := <-sigs

			switch sig {
			case syscall.SIGINT:
				handleSignal(sig)
			case syscall.SIGTERM:
				handleSignal(sig)
				os.Exit(-1)
			case syscall.SIGHUP:
				handleSignal(sig)
			case syscall.SIGUSR1:
				handleSignal(sig)
			default:
				fmt.Println("Ignoring:", sig)
			}
		}
	}()

	for {
		fmt.Println(".")
		time.Sleep(10 * time.Second)
	}
}
