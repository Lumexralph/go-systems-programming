package main

import (
	"fmt"
	"strconv"
)

func main() {
	anArray := [4]int{1, 5, 8, 6}
	aMap := make(map[string]int)
	length := len(anArray)

	for i := 0; i < length; i++ {
		// convert integer to ascii
		fmt.Printf("%s ", strconv.Itoa(i))
		aMap[strconv.Itoa(i)] = anArray[i]
	}

	for key, value := range aMap {
		fmt.Printf("%s: %d\n", key, value)
	}
}
