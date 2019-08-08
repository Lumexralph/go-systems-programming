package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Program executed from process id %d \n", os.Getpid())
}
