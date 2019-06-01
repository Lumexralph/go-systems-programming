package main

import (
	"sort"
	"fmt"
)


type human struct {
	person string
	height int
	weight int
	weightUnit string
}

func main()  {

	users := make([]human, 0)

	researcher := human{"LumexRalph", 180, 72, "kg"}
	teacher := human{"Tosin", 150, 75, "kg"}
	craftsman := human{"Opeyemi", 175, 65, "kg"}
	eventsPlanner := human{"Aanuoluwapo", 180, 72, "kg"}

	users = append(users, researcher, teacher, craftsman, eventsPlanner)

	fmt.Println("0: ", users)

	sort.Slice(users, func(i int, j int) bool {
		return users[i].height < users[j].height
	})

	fmt.Println("Sorting according to height completed: \n", users)

}