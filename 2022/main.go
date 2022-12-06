package main

import (
	"fmt"
	"github.com/farrellit/aoc/2022/rps"
	"os"
	"sort"
)

func main() {
	// Day 1
	if input, err := os.Open("./day1.input1"); err != nil {
		panic(fmt.Errorf("couldn't open day 1 input: %w", err))
	} else {
		elves := ReadCalorieList(input)
		sort.Sort(ElvesByCalories(elves))
		fmt.Println("Day 1:", elves[len(elves)-1].TotalCalories())
		// Day 1 Part 2
		last3 := elves[len(elves)-3:]
		fmt.Println("Day 1 Part 2:", ElvesByCalories(last3).TotalCalories())
		input.Close()
	}
	// Day 2
	if input, err := os.Open("./day2.input1"); err != nil {
		panic(fmt.Errorf("couldn't open day 2 input: %w", err))
	} else {
		game := new(rps.Game)
		if err := game.ReadRounds(input); err != nil {
			panic(fmt.Errorf("Coudln't read rounds from day 2 input: %w", err))
		}
		fmt.Println("Day 2:", game.Scores()[1])
	}
}
