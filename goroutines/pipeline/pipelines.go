package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func numberGenerator(min, max int64, out chan<- int64) {
	var i int64
	for i = min; i <= max; i++ {
		out <- i
	}
	close(out)
}

func findSquares(out chan<- int64, in <-chan int64) {
	for x := range in {
		out <- x * x
	}

	close(out)
}

func calcSum(in <-chan int64) {
	var sum int64
	for x2 := range in {
		sum = sum + x2
	}

	fmt.Printf("The sum of square is %d \n", sum)
}

func main() {
	arguments := os.Args
	if len(arguments) != 3 {
		fmt.Printf("usage: %s n1 n2 \n", filepath.Base(arguments[0]))
		os.Exit(1)
	}

	n1, _ := strconv.ParseInt(arguments[1], 10, 64)
	n2, _ := strconv.ParseInt(arguments[2], 10, 64)
	
	if n1 > n2 {
		fmt.Printf("%d should be smaller than %d \n", n1, n2)
		os.Exit(2)
	}

	naturals := make(chan int64)
	squares := make(chan int64)

	go numberGenerator(n1, n2, naturals)
	go findSquares(squares, naturals)
	calcSum(squares)
}
