package main

import "fmt"

func main() {
	x := -2
	withPointer(&x)
	fmt.Println(x)

	w := newComplex(5, 6)
	fmt.Println(*w)
	fmt.Println(w)
}

func withPointer(x *int) {
	*x = *x * *x
}

type complex struct {
	x, y int
}

func newComplex(x, y int) *complex {
	return &complex{
		x,
		y,
	}
}
