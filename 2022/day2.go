package main

import (
	"fmt"
	"github.com/farrellit/aoc/2022/rps"
	"io"
)

type Day2 struct{}

func (_ Day2) Run(input io.ReadSeeker) (results []interface{}) {
	// Part1
	game := new(rps.Game)
	if err := game.ReadRoundsChoices(input); err != nil {
		panic(fmt.Errorf("Coudln't read rounds from day 2 input: %w", err))
	}
	// Part 2
	input.Seek(0, io.SeekStart)
	game2 := new(rps.Game)
	if err := game2.ReadRounds(
		input, rps.ParseChoice, rps.ParseOutcome); err != nil {
		panic(fmt.Errorf("Couldn't read rounds by choice and desired outcome: %w", err))
	}
	return []interface{}{
		game.Scores()[1],
		game2.Scores()[1],
	}
}
