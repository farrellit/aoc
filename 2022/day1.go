package main

import (
	"io"
	"sort"
)

type Day1 struct{}

func (_ Day1) Run(input io.ReadSeeker) []interface{} {
	elves := ReadCalorieList(input)
	sort.Sort(ElvesByCalories(elves))
	return []interface{}{
		// Part 1
		elves[len(elves)-1].TotalCalories(),
		// Part 2
		ElvesByCalories(elves[len(elves)-3:]).TotalCalories(),
	}
}
