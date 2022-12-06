package main

import (
	"fmt"
	"io"
	"os"
)

// The contract with the daily runnable
type AOCRunnable interface {
	Run(io.ReadSeeker) []interface{}
}

/* represents a day to be run */
type AOCDef struct {
	Solution AOCRunnable
	Input    string
}

type AOC struct {
	Days []AOCDef
}

func (a *AOC) runDayIdx(idx int) {
	if input, err := os.Open(a.Days[idx].Input); err != nil {
		panic(fmt.Errorf("couldn't open day 1 input: %w", err))
	} else {
		defer input.Close()
		for pidx, result := range a.Days[idx].Solution.Run(input) {
			fmt.Printf("Day %2d\tPart %d\t%v\n", idx+1, pidx+1, result)
		}
	}
}

func (a *AOC) Run() {
	for idx := 0; idx < len(a.Days); idx++ {
		a.runDayIdx(idx)
	}
}

func main() {
	aoc := AOC{
		Days: []AOCDef{
			AOCDef{Day1{}, "day1.input1"},
			AOCDef{Day2{}, "day2.input1"},
		},
	}
	aoc.Run()
}
