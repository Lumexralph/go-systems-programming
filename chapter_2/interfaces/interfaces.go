package main

import "fmt"

// I can use interface to achieve a situation where a function
// work with any type and also to do function composition
type coordinates interface {
	xaxis() int
	yaxis() int
}

type point2D struct {
	X int
	Y int
}

// the point2D type has methods that satisfies
// the coordinates interface, what any type that
// wants to be a coordinates should be able to do
func (s point2D) xaxis() int {
	return s.X
}

func (s point2D) yaxis() int {
	return s.Y
}

// a func that takes an interface, I can use it as a gate
// or as a validator in a way to ensure any data or value
// satisfies some criteria using their methods
func findCoordinates(a coordinates) {
	fmt.Printf("X: %d, Y: %d \n", a.xaxis(), a.yaxis())
}

type xpoint int // create a customized type

// add method to the type
func (s xpoint) xaxis() int {
	return int(s)
}

func (s xpoint) yaxis() int {
	return 0
}


func main()  {
	x := point2D{
		X: -1,
		Y: 12,
	}

	fmt.Println(x)
	findCoordinates(x)

	xp := xpoint(10)
	findCoordinates(xp)
}