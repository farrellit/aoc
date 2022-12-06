package main

import (
	"fmt"
	"github.com/farrellit/aoc/2022/rps"
	"io"
)

type Day2 struct{}

func (_ Day2) Run(input io.Reader) []interface{} {
	game := new(rps.Game)
	if err := game.ReadRounds(input); err != nil {
		panic(fmt.Errorf("Coudln't read rounds from day 2 input: %w", err))
	}
	return []interface{}{
		game.Scores()[1],
	}
}
