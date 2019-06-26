package main

import (
	"fmt"
	"time"
)

// https://golang.org/src/time/format.go
func main() {
	t := time.Now()
	fmt.Println("Epoch time: ", t.Unix())
	fmt.Println(t, t.Format(time.RFC3339))
	fmt.Println(t.Weekday(), t.Day(), t.Month(), t.Year())
	time.Sleep(time.Second)
	t1 := time.Now()

	fmt.Println("Time Difference: ", t1.Sub(t))

	formatT := t.Format("13 March 1989")
	fmt.Println(formatT)
	loc, _ := time.LoadLocation("Europe/London")
	londonTime := t.In(loc)
	fmt.Println("London: ", londonTime)

	currentDate := "26 June 2019"
	d, _ := time.Parse("1 August 2000", currentDate)
	fmt.Println(d)

	myDT := "Tuesday 23 May 2017 at 23:36"
	dt, _ := time.Parse("Monday 02 January 2006 at 15:04", myDT)
	fmt.Println(dt)
}
