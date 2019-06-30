package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// define the lower and upper limits of the
// generated random number
func random(min, max int) (value int) {
	value = rand.Intn(max-min) + min
	return
}

func main() {
	var (
		MIN   = 0
		MAX   = 0
		TOTAL = 0
	)

	if len(os.Args) > 3 {
		// convert the string/character to integer
		MIN, _ = strconv.Atoi(os.Args[1])
		MAX, _ = strconv.Atoi(os.Args[2])
		TOTAL, _ = strconv.Atoi(os.Args[3])
	} else {
		fmt.Println("Usage: ", os.Args[0], "MIN MAX TOTAL")
		os.Exit(-1)
	}

	rand.Seed(time.Now().Unix())
	for i := 0; i < TOTAL; i++ {
		generatedNumber := random(MIN, MAX)
		fmt.Print(generatedNumber)
		fmt.Print(" ")
	}
	fmt.Println()
	min := int(^uint(0) >> 1) //largeest 64 bit integer

	fmt.Println(min, int64(^1 << 61))
}
