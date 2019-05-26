package main

import "fmt"


func main() {
	aSlice := []int{-1, 4, 6, 8, 9, 3, 2}
	fmt.Println("Before change: ")
	printSlice(aSlice)
	change(aSlice)
	fmt.Println("After change: ")
	printSlice(aSlice)
}

func change(x []int) {
	x[3] = -2
}

func printSlice(x []int) {
	for _, number := range x {
		fmt.Printf(" %d ", number)
	}
	fmt.Println()
}
